package main

import (
	"context"
	"log"
	"time"

	"github.com/Aavon/playground/grpc_demo/proto"
	"google.golang.org/grpc"
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

	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	res, err := client.Hello(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %s\n", res.Response)

	//stream, err := client.Channel(context.Background())
	//if err != nil {
	//	log.Fatalf("conn err: %v", err)
	//}
	//
	//go func() {
	//	resp := &proto.HelloResponse{}
	//	stream.RecvMsg(resp)
	//	log.Println("resp: ", resp.Response)
	//	time.Sleep(1 * time.Second)
	//	for {
	//		resp := &proto.HelloResponse{}
	//		stream.RecvMsg(resp)
	//		log.Println("resp: ", resp.Response)
	//	}
	//}()

	puller, err := client.Pull(context.Background(), &proto.HelloRequest{})
	if err != nil {
		log.Fatalf("conn err: %v", err)
	}

	go func() {
		resp := &proto.HelloResponse{}
		puller.RecvMsg(resp)
		log.Println("resp: ", resp.Response)
		time.Sleep(10 * time.Second)
		for {
			resp := &proto.HelloResponse{}
			err := puller.RecvMsg(resp)
			log.Println("resp: ", resp.Response)
			if err != nil {
				log.Println("EOF", err)
				break
			}
		}
	}()

	select {}
}
