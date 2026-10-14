package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "productinfo/client/ecommerce/productinfo/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	c, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c2 := pb.NewProductInfoClient(c)
	name := "Google Pixel 9 pro"
	description := `Best of the best phones`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c2.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pid:", r.Value)
	prod, err := c2.GetProduct(ctx, &pb.ProductIDP{Value: r.Value})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(prod.String())
}
