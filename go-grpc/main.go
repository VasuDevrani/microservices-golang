package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if(err != nil) {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()

	// pb.RegisterGreetServiceServer

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatal("Failed to start")
	}
}