package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pd "productinfo/service/ecommerce"

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
	c2 := pd.NewProductInfoClient(c)
	name := "Google Pixel 9 pro"
	description := `Best of the best phones`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c2.AddProduct(ctx, &pd.Product{Name: name, Description: description})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pid:", r.Value)
	prod, err := c2.GetProduct(ctx, &pd.ProductID{Value: r.Value})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(prod.String())
}
