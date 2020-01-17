// +build integration

package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/sysdevguru/checkout-service/api"
)

var (
	testBasketAmount = 10
	basketIDs        = []string{}
)

func TestIntegration(t *testing.T) {
	// create a Basket
	basket, err := grpcClient.CreateBasket(grpcContext, &pb.CreateBasketRequest{})
	assert.Nil(t, err)
	assert.Equal(t, "", basket.Total)
	basketIDs = append(basketIDs, basket.Id)

	// create a bunch of Baskets
	var wg, wg1 sync.WaitGroup
	for i := 0; i < testBasketAmount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b, err := grpcClient.CreateBasket(grpcContext, &pb.CreateBasketRequest{})
			assert.Equal(t, nil, err)
			assert.Equal(t, "", b.Total)
			basketIDs = append(basketIDs, b.Id)
		}()
		wg.Wait()
	}

	// add products to first basket
	products := []string{"PEN", "PEN", "PEN", "TSHIRT", "TSHIRT", "TSHIRT", "MUG", "PEN", "MUG"}
	for _, v := range products {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			resp, err := grpcClient.AddToBasket(grpcContext, &pb.AddProductRequest{Basketid: basketIDs[0], Product: v})
			assert.Nil(t, err)
			assert.Equal(t, 200, int(resp.Code))
			assert.Equal(t, "", resp.Error)
		}()
		wg1.Wait()
	}

	// get first basket amount
	resp, err := grpcClient.GetBasketAmount(grpcContext, &pb.GetAmountRequest{Basketid: basketIDs[0]})
	assert.Nil(t, err)
	assert.Equal(t, "70.00€", resp.Total)

	// remove Basket
	resp1, err := grpcClient.RemoveBasket(grpcContext, &pb.RemoveBasketRequest{Basketid: basketIDs[0]})
	assert.Nil(t, err)
	assert.Equal(t, 200, int(resp1.Code))
	assert.Equal(t, "", resp1.Error)
}
