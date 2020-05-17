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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", req)
	number := req.GetNumber()
	k := int32(2)
	for number > 1 {
		if number%k == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{Result: k}
			stream.Send(res)
			number = number / k
		} else {
			k++
		}
	}
	return nil
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
