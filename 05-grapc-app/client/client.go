package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

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
	doServerStreaming(ctx, client)
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
	for {
		response, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println("Prime No : ", response.GetPrimeNo())
	}
}
