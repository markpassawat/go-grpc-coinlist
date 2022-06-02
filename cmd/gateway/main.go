package main

import (
	// "context"
	// "io"
	// "log"
	// "time"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/markpassawat/go-grpc-coinlist/cmd/gateway/config"
	"github.com/markpassawat/go-grpc-coinlist/cmd/gateway/server"
	"github.com/kelseyhightower/envconfig"

)


func main() {
	cfg := config.Config{}
	log := logrus.StandardLogger()
	envconfig.MustProcess("GRPC_GATEWAY", &cfg)

	log.Info("starting server...")
	e := server.Handler(&cfg)
	logrus.Info(fmt.Sprintf("Serving gRPC-Gateway on %s", e.Addr))
	logrus.Fatalln(e.ListenAndServe())
}
