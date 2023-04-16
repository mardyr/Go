package main

import (
	"context"
	"log"

	pb "github.com/mardyr/Go/gRPC/greet/proto"
)

func doGreet(cc pb.GreetServiceClient) {
	log.Println(" doGreet invoked ")
	res, err := cc.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v \n", err)
	}

	log.Printf("Greeting: %s \n", res.Result)
}
