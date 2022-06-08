package server

import (
	"context"
	"fmt"

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

	res, err := db.GetCoinById(in.CoinId)

	errTemp := fmt.Errorf("There is no coin named %v", in.CoinId)

	if err != nil {
		return res, errTemp
	} else {
		return res, nil
	}
}

func (s *coinServer) CreateCoins(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}

	isExist, errIsExist := db.IsExist(in.CoinId)
	var err error

	if errIsExist != nil {
		err = fmt.Errorf("There is something wrong with database")
	}
	if isExist {
		err = fmt.Errorf("%v is already exist", in.CoinId)
	} else {
		isCreated := db.InsertOne(in.CoinId)
		if isCreated {
			res.Status = 201
		} else {
			err = fmt.Errorf("There is no coin named %v", in.CoinId)
		}
	}

	return res, err
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
