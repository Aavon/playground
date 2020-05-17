package main

import (
	"context"
	"github.com/Aavon/playground/grpc_demo/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewDemoClient(conn)
	req := &proto.HelloRequest{
		Say: "Aaron",
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %s\n", res.Response)
}
