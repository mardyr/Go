package main

import (
	"log"

	pb "github.com/mardyr/Go/gRPC/calculator/proto"
)

func (*Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Decomposite primes values from %v \n", in)

	primeStart := int64(2)
	valueReceived := in.Number

	for valueReceived > 1 {
		if valueReceived%primeStart == 0 {
			log.Printf("Prime value %d send \n", primeStart)
			stream.Send(&pb.PrimeResponse{
				Result: primeStart,
			})
			valueReceived /= primeStart
		} else {
			primeStart++
		}

	}
	return nil
}
