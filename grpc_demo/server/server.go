package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Aavon/playground/grpc_demo/proto"
	"google.golang.org/grpc"
)

type MyDemo struct{}

func (h *MyDemo) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Say:%s\n", req.Say)
	res := &proto.HelloResponse{
		Response: fmt.Sprintf("Hello %s", req.Say),
	}
	return res, nil
}

func main() {
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterDemoServer(grpcServer, &MyDemo{})
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
