package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := grpc.NewServer()
	pb.regprodinf(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
