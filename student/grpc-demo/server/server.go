package main

import (
	"context"
	"google.golang.org/grpc"
	"java-to-go-learning/student/grpc-demo/proto"
	"log"
	"net"
)

// server用于实现hello.HelloServiceServer
type server struct {
	proto.UnimplementedHelloServiceServer
}

// SayHello实现hello.HelloServiceServer接口
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &proto.HelloReply{Response: "Hello " + in.GetGreeting()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
