package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	
	"github.com/markpassawat/go-grpc-coinlist/cmd/coin_list/config"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

func Handler(cfg *config.Config) (*grpc.Server, net.Listener) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterCoinListServer(s, &coinServer{})
	return s, lis
}

