package main

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

type server struct {
	pM map[string]*pb.Product
}

func (s *server) add(ctx context.Context, in *pd.prod) (*pb.pid, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error", err)
	}
	in.Id = out.String()
	if s.pM == nil {
		s.pM == make(map[string]*pb.prod)
	}
	s.pM[in.Id] = in
	return &pd.pid{Value: in.Id}, status.New(codes.OK, "").Err()
}
