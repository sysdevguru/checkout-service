package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sysdevguru/checkout-service/api"
	"google.golang.org/grpc"
)

var (
	grpcClient  pb.BasketServiceClient
	grpcContext context.Context
)

func init() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("checkout-client: did not connect to grpc server: %v", err)
	}
	grpcClient = pb.NewBasketServiceClient(conn)
	grpcContext, _ = context.WithTimeout(context.Background(), time.Second)
}

func main() {
	// create a basket
	basket, err := grpcClient.CreateBasket(grpcContext, &pb.CreateBasketRequest{})
	if err != nil {
		log.Fatalf("checkout-client: creating a basket failure: %v", err)
	}
	log.Printf("checkout-client: create a basket: %v\n", basket.Id)

	// add products to basket
	products := []string{"PEN", "TSHIRT", "PEN", "PEN", "MUG", "TSHIRT"}
	for _, v := range products {
		_, err = grpcClient.AddToBasket(grpcContext, &pb.AddProductRequest{
			Basketid: basket.Id,
			Product:  v,
		})
		if err != nil {
			log.Fatalf("checkout-client: adding product to basket failure: %v", err)
		}
		log.Printf("checkout-client: added product %v to basket %v\n", v, basket.Id)
	}

	// get total of the basket
	resp, err := grpcClient.GetBasketAmount(grpcContext, &pb.GetAmountRequest{
		Basketid: basket.Id,
	})
	if err != nil {
		log.Fatalf("checkout-client: getting amount of basket failure: %v", err)
	}
	log.Printf("checkout-client: basket %v amount is %v\n", basket.Id, resp.Total)

	// remove basket
	_, err = grpcClient.RemoveBasket(grpcContext, &pb.RemoveBasketRequest{
		Basketid: basket.Id,
	})
	if err != nil {
		log.Fatalf("checkout-client: removing basket failure: %v", err)
	}
	log.Printf("checkout-client: removed basket %v\n", basket.Id)
}
