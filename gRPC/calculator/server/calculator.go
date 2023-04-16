package main

import (
	"context"
	"log"

	pb "github.com/mardyr/Go/gRPC/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {

	log.Printf("Calculate sum for values %v \n", in)

	return &pb.CalculatorResponse{
		Result: in.FirstValue + in.SecondValue,
	}, nil
}
