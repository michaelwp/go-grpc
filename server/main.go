package main

import (
	"github.com/michaelwp/go-gRpc/server/chat"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen n port 9000 %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRpc server over port 9000: %v", err)
	}
}
