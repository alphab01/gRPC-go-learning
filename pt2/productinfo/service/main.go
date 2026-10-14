package main

import (
	"log"
	"net"
	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Print("start")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
