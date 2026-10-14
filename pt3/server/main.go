package main

import "github.com/golang/protobuf/ptypes/wrappers"

func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrderServer) error {

}
