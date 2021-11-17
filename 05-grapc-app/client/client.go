package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	/* request & response */
	//doRequestResponse(ctx, client)

	/* server streaming */
	//doServerStreaming(ctx, client)

	/* client streaming */
	//doClientStreaming(ctx, client)

	/* bidirectional streaming */
	doBidiStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	request := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	response, err := client.Add(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.GetSum())
}

/* Server Streaming */
func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	request := &proto.PrimeRequest{
		Start: 10,
		End:   100,
	}
	stream, err := client.GeneratePrime(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Delaying by 10 Secs")
	time.Sleep(10 * time.Second)
	for {
		response, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println("Prime No : ", response.GetPrimeNo())
	}
}

/* Client Streaming */
func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	nos := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending : ", no)
		req := &proto.AverageRequest{
			Num: no,
		}
		stream.Send(req)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average : ", response.GetAverage())
}

func doBidiStreaming(ctx context.Context, client proto.AppServiceClient) {
	users := []proto.User{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Rajan"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}
	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		for _, user := range users {
			req := &proto.GreetRequest{
				User: &user,
			}

			time.Sleep(5 * time.Second)
			log.Println("Sending : ", fmt.Sprintf("%v", user))
			stream.Send(req)
		}
		log.Println("Sent all the requests")
	}()
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Received all responses")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Message : ", res.GetMessage())
		}
		done <- true
	}()
	<-done
}
