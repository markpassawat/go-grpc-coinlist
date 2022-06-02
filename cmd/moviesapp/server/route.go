package server

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	pb "github.com/markpassawat/go-grpc-coinlist/proto/moviesapp"
)

type movieServer struct {
	pb.UnimplementedMovieServer
}

var movies []*pb.MovieInfo

// func NewServer() *movieServer {
// 	return &movieServer{}
// }

func (s *movieServer) GetMovies(in  *pb.Empty, stream pb.Movie_GetMoviesServer) error {
	for _, movie := range movies {
		if err := stream.Send(movie); err != nil {
			return err
		}
	}
	return nil
}

func (s *movieServer) GetMovie(ctx context.Context,
	in *pb.Id) (*pb.MovieInfo, error) {

		res := &pb.MovieInfo{}

		for _, movie := range movies {
			if movie.GetId() == in.GetValue() {
				res = movie
				break
			}
		}
	
		return res, nil
	}

func (s *movieServer) CreateMovie(ctx context.Context,
	in *pb.MovieInfo) (*pb.Id, error) {
	res := &pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000000))
	in.Id = res.GetValue()
	movies = append(movies, in)
	log.Printf("Received: %v", movies)

	return res, nil

}

func (s *movieServer) UpdateMovie(ctx context.Context,
	in *pb.MovieInfo) (*pb.Status, error) {

	res := &pb.Status{}
	for index, movie := range movies {
		if movie.GetId() == in.GetId() {
			movies = append(movies[:index], movies[index+1:]...)
			in.Id = movie.GetId()
			movies = append(movies, in)
			res.Value = 1
			break
		}
	}

	return res, nil
}

func (s *movieServer) DeleteMovie(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {

	res := &pb.Status{}
	for index, movie := range movies {
		if movie.GetId() == in.GetValue() {
			movies = append(movies[:index], movies[index+1:]...)
			res.Value = 1
			break
		}
	}

	return res, nil
}
