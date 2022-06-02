package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	
	"github.com/markpassawat/go-grpc-coinlist/cmd/moviesapp/config"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/moviesapp"
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
	pb.RegisterMovieServer(s, &movieServer{})
	return s, lis
}

// func main() {
// 	initMovies()
// 	s := grpc.NewServer()

// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalln("Failed to listen:", err)
// 	}

// 	pb.RegisterMovieServer(s, &movieServer{})

// 	log.Printf("gRPC server listening on %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

// func initMovies() {
// 	movie1 := &pb.MovieInfo{Id: "1", Isbn: "0593310438",
// 		Title: "The Batman", Director: &pb.Director{
// 			Firstname: "Matt", Lastname: "Reeves"}}
// 	movie2 := &pb.MovieInfo{Id: "2", Isbn: "3430220302",
// 		Title: "Doctor Strange in the Multiverse of Madness",
// 		Director: &pb.Director{Firstname: "Sam",
// 			Lastname: "Raimi"}}

// 	movies = append(movies, movie1)
// 	movies = append(movies, movie2)
// }

