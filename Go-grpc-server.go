package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/razeencheng/demo-go/grpc/demo2/helloworld"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	_ "google.golang.org/grpc"
)

type SayHelloServer struct{}

func (s *SayHelloServer) SayHelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (res *pb.HelloWorldResponse, err error) {
	log.Printf("Client Greeting:%s", in.Greeting)
	log.Printf("Client Info:%v", in.Infos)

	var an *any.Any
	if in.Infos["hello"] == "world" {
		an, err = ptypes.MarshalAny(&pb.HelloWorld{Msg: "Good Request"})
	} else {
		an, err = ptypes.MarshalAny(&pb.Error{Msg: []string{"Bad Request", "Wrong Info Msg"}})
	}

	if err != nil {
		return
	}
	return &pb.HelloWorldResponse{
		Reply:   "Hello World !!",
		Details: []*any.Any{an},
	}, nil
}
func main() {

	// 我们首先须监听一个tcp端口
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	// 新建一个grpc服务器
	grpcServer := grpc.NewServer()
	// 向grpc服务器注册SayHelloServer
	pb.RegisterHelloWorldServiceServer(grpcServer, &SayHelloServer{})
	// 启动服务
	grpcServer.Serve(lis)
}