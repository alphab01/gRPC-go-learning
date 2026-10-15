package main

import "log"

func main() {
	orderMgtClient := pb.NewOrderManagementClient(conn)
	retrievedOrder, err := orderMgtClient.GetOrder(ctx, &wrapper.StringValue{Value: "106"})
	log.Print(retrievedOrder)
}
