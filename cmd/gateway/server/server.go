package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/markpassawat/go-grpc-coinlist/cmd/gateway/config"
	"github.com/markpassawat/go-grpc-coinlist/pkg/middleware"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

func Handler(cfg *config.Config) *http.Server {
	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("Config: %+v", cfg)
	}
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.Coinlist,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterCoinListHandler(context.Background(), gwmux, conn)
	if err != nil {
		logrus.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: middleware.Cors(&middleware.CorsConfig{
			AllowedOrigin: cfg.AllowedOrigin,
		})(gwmux),
	}
	return gwServer
}
