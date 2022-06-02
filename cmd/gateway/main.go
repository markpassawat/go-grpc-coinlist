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

// const (
// 	address = "localhost:50051"
// )
func main() {
	cfg := config.Config{}
	log := logrus.StandardLogger()
	envconfig.MustProcess("GRPC_GATEWAY", &cfg)

	log.Info("starting server...")
	e := server.Handler(&cfg)
	logrus.Info(fmt.Sprintf("Serving gRPC-Gateway on %s", e.Addr))
	logrus.Fatalln(e.ListenAndServe())
}
// func main() {

// 	conn, err := grpc.Dial(address, grpc.WithInsecure(),
// 		grpc.WithBlock())

// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	// client := pb.NewMovieClient(conn)

// 	gwmux := runtime.NewServeMux()

// 	err = pb.RegisterMovieHandler(context.Background(), gwmux, conn)
// 	if err != nil {
// 		logrus.Fatalln("Failed to register gateway:", err)
// 	}

// 	e, lis := server.Handler(&cfg)

// 	// runGetMovie(client, "1")
// 	// runCreateMovie(client, "24325645", "Spiderman Spiderverse",
// 	// 	"Stan", "Lee")
// 	// runUpdateMovie(client, "98498081", "24325645", "Spiderman Spiderverse",
// 	// 	"Peter", "Parker")
// 	// runDeleteMovie(client, "98498081")
// }

// func runGetMovies(client pb.MovieClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.Empty{}
// 	stream, err := client.GetMovies(ctx, req)
// 	if err != nil {
// 		log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
// 	}
// 	for {
// 		row, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
// 		}
// 		log.Printf("MovieInfo: %v", row)
// 	}
// }

// func runGetMovie(client pb.MovieClient, movieid string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.Id{Value: movieid}
// 	res, err := client.GetMovie(ctx, req)
// 	if err != nil {
// 		log.Fatalf("%v.GetMovie(_) = _, %v", client, err)
// 	}
// 	log.Printf("MovieInfo: %v", res)
// }

// func runCreateMovie(client pb.MovieClient, isbn string,
// 	title string, firstname string, lastname string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.MovieInfo{Isbn: isbn, Title: title,
// 		Director: &pb.Director{Firstname: firstname,
// 			Lastname: lastname}}
// 	res, err := client.CreateMovie(ctx, req)
// 	if err != nil {
// 		log.Fatalf("%v.CreateMovie(_) = _, %v", client, err)
// 	}
// 	if res.GetValue() != "" {
// 		log.Printf("CreateMovie Id: %v", res)
// 	} else {
// 		log.Printf("CreateMovie Failed")
// 	}

// }

// func runUpdateMovie(client pb.MovieClient, movieid string,
// 	isbn string, title string, firstname string, lastname string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.MovieInfo{Id: movieid, Isbn: isbn,
// 		Title: title, Director: &pb.Director{
// 			Firstname: firstname, Lastname: lastname}}
// 	res, err := client.UpdateMovie(ctx, req)
// 	if err != nil {
// 		log.Fatalf("%v.UpdateMovie(_) = _, %v", client, err)
// 	}
// 	if int(res.GetValue()) == 1 {
// 		log.Printf("UpdateMovie Success")
// 	} else {
// 		log.Printf("UpdateMovie Failed")
// 	}
// }

// func runDeleteMovie(client pb.MovieClient, movieid string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.Id{Value: movieid}
// 	res, err := client.DeleteMovie(ctx, req)
// 	if err != nil {
// 		log.Fatalf("%v.DeleteMovie(_) = _, %v", client, err)
// 	}
// 	if int(res.GetValue()) == 1 {
// 		log.Printf("DeleteMovie Success")
// 	} else {
// 		log.Printf("DeleteMovie Failed")
// 	}
// }
