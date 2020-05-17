package main

import (
	"context"
	"fmt"
	"github.com/go_udemy/grpc-go-course/calculator/calculatorpb"
	"github.com/go_udemy/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)
	doUnary(c)

	l := calculatorpb.NewCalculatorServiceClient(cc)
	doUnaryC(l)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "George",
			LastName:  "Georgakas",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doUnaryC(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Calculator Unary RPC")
	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			NumberOne: 10,
			NumberTwo: 3,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Response from Calculator: %v", res.Result)
}
