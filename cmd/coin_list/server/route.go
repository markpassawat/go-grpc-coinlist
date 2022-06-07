package server

import (
	"context"
	"fmt"
	"log"

	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

type coinServer struct {
	pb.UnimplementedCoinListServer
}

var coins []*pb.CoinInfo

func (s *coinServer) GetCoins(in *pb.Empty, stream pb.CoinList_GetCoinsServer) error {

	coinlist := db.GetAllCoin()

	for _, coin := range coinlist {
		if err := stream.Send(coin); err != nil {
			return err
		}
	}
	return nil

}

func (s *coinServer) GetCoin(ctx context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	// res := &pb.CoinInfo{}

	res, err := db.GetCoinById(in.CoinId)

	if err != nil {
		// log.Fatal("asd")
		fmt.Print(err)
		return res, nil
	} else {
		return res, nil
	}
}

func (s *coinServer) CreateCoins(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}

	isExist, err := db.IsExist(in.CoinId)

	if err != nil {
		log.Fatal("ERR: ", err)
		res.Status = 400
	}
	if isExist {
		res.Status = 400
	} else {
		isCreated := db.InsertOne(in.CoinId)
		if isCreated {
			res.Status = 201
		} else {
			res.Status = 400
		}
	}

	return res, nil
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

func (s *coinServer) SearchCoins(in *pb.InputText, stream pb.CoinList_SearchCoinsServer) error {

	coinlist := db.SearchCoins(in.InputText)

	for _, coin := range coinlist {
		if err := stream.Send(coin); err != nil {
			return err
		}
	}

	return nil
}
