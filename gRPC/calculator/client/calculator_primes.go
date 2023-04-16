package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mardyr/Go/gRPC/calculator/proto"
)

func takePrimes(csc pb.CalculatorServiceClient) {

	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{
		Number: 120,
	}
	stream, err := csc.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Result)
	}
}
