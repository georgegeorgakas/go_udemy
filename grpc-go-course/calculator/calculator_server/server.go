package main

import (
	"context"
	"fmt"
	"github.com/go_udemy/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator funcation was invoked with %v", req)

	numberOne := req.GetNumberOne()
	numberTwo := req.GetNumberTwo()
	result := numberOne + numberTwo
	res := &calculatorpb.SumResponse{Result: result}

	return res, nil
}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to liste: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
