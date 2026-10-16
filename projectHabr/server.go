package main

type aserver {
  pb.UnimplementedAuthServer
  postpuul *pgxpool.Pool
  redcli *redis.Client
  jwtsec string
}

func newstr() *aserver {
  return &aserver{}
}
