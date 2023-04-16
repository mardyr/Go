package main

import (
	"context"
	"log"

	pb "github.com/mardyr/Go/gRPC/calculator/proto"
)

func makeCalculation(csc pb.CalculatorServiceClient) {

	var firstVal int32 = 3
	var secondVal int32 = 10

	log.Printf("Send request to calculate sum for: %d and %d \n", firstVal, secondVal)

	res, err := csc.Calculate(context.Background(), &pb.CalculatorRequest{
		FirstValue:  firstVal,
		SecondValue: secondVal,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v \n", err)
	}

	log.Printf("Sum is equal to: %v \n", res.Result)
}
