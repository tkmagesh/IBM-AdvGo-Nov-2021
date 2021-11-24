package main

import (
	"context"
	"grpc-chat/proto"
	"log"
	"net"
	"os"
	"sync"

	grpc "google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Connection struct {
	stream   proto.Broadcast_CreateStreamServer
	id       string
	active   bool
	userName string
	error    chan error
}

//stateful server
type Server struct {
	Connection []*Connection
	proto.UnimplementedBroadcastServer
}

func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream:   stream,
		id:       pconn.User.Id,
		userName: pconn.GetUser().GetName(),
		active:   true,
		error:    make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	//done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *proto.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("Sending message to: ", conn.stream)

				if err != nil {
					grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	/* go func() {
		wait.Wait()
		close(done)
	}() */
	wait.Wait()
	//<-done
	return &proto.Close{}, nil
}

func main() {
	var connections []*Connection

	server := &Server{Connection: connections}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting server at port :8085")

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
}