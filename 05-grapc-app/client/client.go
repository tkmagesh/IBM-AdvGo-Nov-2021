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
