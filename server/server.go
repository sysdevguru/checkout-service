package server

import (
	"context"
	"log"
	"os/exec"

	pb "github.com/sysdevguru/checkout-service/api"
	"github.com/sysdevguru/checkout-service/common"
	"github.com/sysdevguru/checkout-service/util"
)

var (
	basketMap  map[string]pb.Basket
)

type GRPCServer struct{}

func init() {
	// initialize database
	basketMap = make(map[string]pb.Basket)
}

// CreateBasket creates Basket and returns the created Basket
func (s *GRPCServer) CreateBasket(ctx context.Context, in *pb.CreateBasketRequest) (*pb.Basket, error) {
	basketID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Printf("checkout-server: unexpected basket_id generation failure: %v\n", err)
		return &pb.Basket{}, err
	}
	basket := pb.Basket{
		Id: string(basketID),
		Items: []string{},
		Total: "",
	}

	// add to db
	basketMap[string(basketID)] = basket
	log.Printf("checkout-server: created a basket %v\n", basket.Id)
	return &basket, nil
}

// AddToBasket adds product to Basket
func (s *GRPCServer) AddToBasket(ctx context.Context, in *pb.AddProductRequest) (*pb.BasketResponse, error) {
	basketID := in.GetBasketid()
	product := in.GetProduct()
	basket := basketMap[basketID]
	basket.Items = append(basket.Items, product)
	basketMap[basketID] = basket

	log.Printf("checkout-server: added a product %v to basket %v\n", product, basketID)

	return &pb.BasketResponse{
		Code:  200,
		Error: "",
	}, nil
}

// GetBasketAmount returns total amount of basket
func (s *GRPCServer) GetBasketAmount(ctx context.Context, in *pb.GetAmountRequest) (*pb.GetAmountResponse, error) {
	basketID := in.GetBasketid()
	items := basketMap[basketID].Items
	total := util.CalculateAmount(items)
	basket := basketMap[basketID]
	basket.Total = total + common.Server.Currency

	log.Printf("checkout-server: Total amount %v%v of basket %v\n", total, common.Server.Currency, basketID)
	
	return &pb.GetAmountResponse{
		Total: total + common.Server.Currency,
	}, nil
}

// RemoveBasket removes basket
func (s *GRPCServer) RemoveBasket(ctx context.Context, in *pb.RemoveBasketRequest) (*pb.BasketResponse, error) {
	basketID := in.GetBasketid()
	delete(basketMap, basketID)

	log.Printf("checkout-server: removed a basket %v\n", basketID)

	return &pb.BasketResponse{
		Code:  200,
		Error: "",
	}, nil
}
