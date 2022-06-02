package main

import (
	"github.com/markpassawat/go-grpc-coinlist/cmd/moviesapp/config"
	"github.com/markpassawat/go-grpc-coinlist/cmd/moviesapp/server"
	
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	log := logrus.StandardLogger()
	envconfig.MustProcess("GRPC_MOVIESAPP", &cfg)

	log.Info("starting server...")
	e, lis := server.Handler(&cfg)
	// Serve gRPC Server
	log.Println("Serving gRPC on", lis.Addr().String())
	log.Fatal(e.Serve(lis))
}