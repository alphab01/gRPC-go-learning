package main

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type aserver struct {
	pb.UnimplementedAuthServer
	postpool *pgxpool.Pool
	redcli   *redis.Client
	jwtsec   string
}

func newstr() *aserver {
	return &aserver{}
}

func stserv(ctx context.Context, pool *pgxpool.Pool, redis *redis.Client, jwt string) {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	aserv := &internal.aserver{
		postpool: pool,
		redcli:   redis,
		jwtsec:   jwt,
	}
	grpcserv := grpc.NewServer()
	pb.RegAuthServer(grpcserv, aserv)
	go func() {
		log.Porintln("begin")
		if err := grpcserv.Serve(l); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()
	log.Println("get signal")
	grpcserv.GracefulStop()
	log.Println("server end")
}
