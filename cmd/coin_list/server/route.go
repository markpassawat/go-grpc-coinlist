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

func (s *coinServer) GetCoins(ctx context.Context, in *pb.Empty) (*pb.ReturnList, error) {

	coinList := db.GetAllCoin()

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil

}

func (s *coinServer) GetCoin(ctx context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	res, err := db.GetCoinById(in.CoinId)

	errAlert := fmt.Errorf("There is no coin named %v", in.CoinId)

	if err != nil {
		return res, errAlert
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

func (s *coinServer) SearchCoins(ctx context.Context, in *pb.SearchText) (*pb.ReturnList, error) {

	coinList := db.SearchCoins(in.InputText)

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil

}
