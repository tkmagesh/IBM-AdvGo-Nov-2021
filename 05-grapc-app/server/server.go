package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Println("Add request received for ", x, y)
	response := &proto.AddResponse{
		Sum: result,
	}
	return response, nil
}

func (s *server) GeneratePrime(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimeServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for i := start; i <= end; i++ {
		if isPrime(i) {
			time.Sleep(500 * time.Millisecond)
			res := &proto.PrimeResponse{
				PrimeNo: i,
			}
			stream.Send(res)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	if no == 1 {
		return false
	}
	if no == 2 {
		return true
	}
	for i := int32(2); i < int32(no); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) CalculateAverage(stream proto.AppService_CalculateAverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := sum / count
			res := &proto.AverageResponse{
				Average: average,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		num := req.GetNum()
		sum += num
		count++
	}
	return nil
}

func main() {
	s := &server{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}
