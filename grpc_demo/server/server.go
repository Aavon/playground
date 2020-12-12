package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Aavon/playground/grpc_demo/proto"
	"google.golang.org/grpc"
)

type MyDemo struct{}

func (h *MyDemo) Channel(conn proto.Demo_ChannelServer) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count := 0
		defer wg.Done()
		for {
			//msg := &proto.HelloRequest{}
			//err := conn.RecvMsg(msg)
			//if err != nil {
			//	log.Println("recv err: ", err)
			//	return
			//}
			//fmt.Println("recv: ", msg.Say)
			//resp := &proto.HelloResponse{}
			//resp.Response = fmt.Sprintf("%s received.", msg.Say)
			//err = conn.SendMsg(resp)
			//if err != nil {
			//	log.Println("send err: ", err)
			//	return
			//}
			for i := 0; i < 100; i++ {
				if count == 14727 {
					log.Println("break")
					return
				}
				resp := &proto.HelloResponse{}
				resp.Response = fmt.Sprintf("%d", i)
				err := conn.SendMsg(resp)
				if err != nil {
					log.Println("send err: ", err)
					return
				}
				count++
				log.Println(count)
			}
		}
	}()

	wg.Wait()
	return nil
}

func (h *MyDemo) Pull(req *proto.HelloRequest, s proto.Demo_PullServer) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count := 0
		defer wg.Done()
		for {
			for i := 0; i < 100; i++ {
				//if count == 54727 {
				//	log.Println("break")
				//	return
				//}
				resp := &proto.HelloResponse{}
				resp.Response = fmt.Sprintf("%d", count)
				err := s.SendMsg(resp)
				if err != nil {
					log.Println("send err: ", err)
					return
				}
				count++
				log.Println(count)
			}
		}
	}()

	wg.Wait()
	log.Println("exit.")
	return nil
}

func (h *MyDemo) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Say:%s\n", req.Say)
	res := &proto.HelloResponse{
		Response: fmt.Sprintf("Hello %s", req.Say),
	}
	//time.Sleep(10 * time.Second)
	return res, nil
}

func main() {
	//filter := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	log.Println("filter: ", info)
	//	return &proto.HelloResponse{Response: "filtered"}, nil
	//	//return handler(ctx, req)
	//}
	//
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))
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
