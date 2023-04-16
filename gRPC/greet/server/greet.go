package main

import (
	"context"
	"log"

	pb "github.com/mardyr/Go/gRPC/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v \n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil

}
