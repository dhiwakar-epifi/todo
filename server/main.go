package main

import (
	"log"
	"net"

	pb "example.com/grpc-todo/proto"
	"google.golang.org/grpc"

	wire "example.com/grpc-todo/wire"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	fireflySvc := wire.InitialiseTodoService()

	pb.RegisterTodoServiceServer(s, fireflySvc)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
