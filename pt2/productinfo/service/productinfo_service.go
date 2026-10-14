package main

import (
	"context"

	pb "productinfo/service/ecommerce"
	pd "productinfo/service/ecommerce"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pd.Product
}

func (s *server) AddProduct(ctx context.Context, in *pd.Product) (*pd.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "err", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pd.Product)
	}
	s.productMap[in.Id] = in
	return &pd.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *pd.ProductID) (*pd.Product, error) {
	val, exists := s.productMap[in.Value]
	if exists {
		return val, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "err", in.Value)
}
