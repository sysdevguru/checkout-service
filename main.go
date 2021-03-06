package main

import (
	"log"
	"net"

	pb "github.com/sysdevguru/checkout-service/api"
	"github.com/sysdevguru/checkout-service/common"
	"github.com/sysdevguru/checkout-service/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", common.Server.GRPCPort)
	if err != nil {
		log.Printf("checkout-server: failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterBasketServiceServer(s, &server.GRPCServer{})
	if err := s.Serve(lis); err != nil {
		log.Printf("checkout-server: failed to serve: %v\n", err)
	}
}
