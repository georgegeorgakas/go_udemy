package main

import (
	"fmt"
	"ggeorgakas/grpc-go-course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {}

func main() {
	fmt.Printf("Greet server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to liste: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
