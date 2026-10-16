package main

import (
  "log"
)

type aserver {
  pb.UnimplementedAuthServer
  postpool *pgxpool.Pool
  redcli *redis.Client
  jwtsec string
}

func newstr() *aserver {
  return &aserver{}
}

func stserv(ctx context.Context, pool *pgxpool.Pool, redis *redis.Client, jwt string) {
  l, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatal(err)
  }
  aserv := &internal.aserver {
    postpool :pool,
    redcli: redis,
    jwtsec: jwt,
  }
}
