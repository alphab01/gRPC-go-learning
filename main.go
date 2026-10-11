package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func (s *server) add(ctx context.Context, in *pd.prod) (*pb.prod, error) {

}

func (s *server) get(ctx context.Context, in *pd.pid) (*pb.pid, error) {

}

func main() {
	l, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pd.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("err", err)
	}
}
