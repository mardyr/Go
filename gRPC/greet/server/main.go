package main

import (
	"log"
	"net"

	pb "github.com/mardyr/Go/gRPC/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to start listening on :%v \n", err)
	}

	log.Printf("Listening on %s address \n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s,&Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v \n", err)
	}

}
