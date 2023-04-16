package main

import (
	"log"
	"net"

	pb "github.com/mardyr/Go/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var serverAddr string = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", serverAddr)

	if err != nil {
		log.Fatalf("Failed to start listening on :%v \n", err)
	}

	log.Printf("Listening on %s address \n", serverAddr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v \n", err)
	}

}
