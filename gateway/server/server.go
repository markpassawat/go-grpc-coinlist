package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/markpassawat/go-gprc-coinlist/gateway/config"
	// "github.com/markpassawat/gateway/pkg/middleware"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/moviesapp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		cfg.Helloworld,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterMovieHandler(context.Background(), gwmux, conn)
	if err != nil {
		logrus.Fatalln("Failed to register gateway:", err)
	}

	// gwServer := &http.Server{
	// 	Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	// 	Handler: middleware.Cors(&middleware.CorsConfig{
	// 		AllowedOrigin: cfg.AllowedOrigin,
	// 	})(gwmux),
	// }
	// return gwServer
}