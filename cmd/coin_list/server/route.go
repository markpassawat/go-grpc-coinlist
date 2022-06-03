package server

import (
	"context"
	"log"
	// "log"
	// "math/rand"
	// "strconv"

	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

type coinServer struct {
	pb.UnimplementedCoinListServer
}

var coins []*pb.CoinInfo

func (s *coinServer) GetCoins(in *pb.Empty, stream pb.CoinList_GetCoinsServer) error {
	for _, coin := range coins {
		if err := stream.Send(coin); err != nil {
			return err
		}
	}
	return nil
}

func (s *coinServer) GetCoin(ctx context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {
	
		res := &pb.CoinInfo{}

	for _, coin := range coins {
		if coin.GetSymbol() == in.GetSymbol() {
			res = coin
			break
		}
	}

	return res, nil
}

func (s *coinServer) CreateCoins(ctx context.Context,
	in *pb.CoinInfo) (*pb.Id, error) {
	res := pb.Id{}
	res.CoinId = in.CoinId
	res.Symbol = in.Symbol
	res.Name = in.Name

	coins = append(coins, in)
	return &res,nil
}

func (s *coinServer) UpdateCoins(ctx context.Context,
	in *pb.CoinInfo) (*pb.Status, error) {

	res := &pb.Status{}
	for index, coin := range coins {
		if coin.GetCoinId() == in.GetCoinId() {
			coins = append(coins[:index], coins[index+1:]...)
			in.CoinId = coin.GetCoinId()
			coins = append(coins, in)
			res.Status = 200
			break
		}
	}

	return res, nil
}

func (s *coinServer) DeleteCoin(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}
	res.Status = 204
	for index, coin := range coins {
		if coin.GetSymbol() == in.GetSymbol() {
			coins = append(coins[:index], coins[index+1:]...)
			res.Status = 200
			break
		}
	}

	return res, nil
}
