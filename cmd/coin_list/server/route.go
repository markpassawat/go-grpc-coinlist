package server

import (
	"context"
	"fmt"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	bun "github.com/uptrace/bun"
)

type CoinServer struct {
	pb.UnimplementedCoinListServer
	db *bun.DB
}

var coins []*pb.CoinInfo

func (s *CoinServer) GetCoins(ctxCon context.Context, in *pb.Empty) (*pb.ReturnList, error) {
	// TestConnection(s.db)
	coinList := db.GetAllCoin(s.db, ctxCon)

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil
}

func (s *CoinServer) GetCoin(ctxCon context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	res, err := db.GetCoinById(s.db, ctxCon, in.CoinId)

	errAlert := fmt.Errorf("There is no coin named %v", in.CoinId)

	if err != nil {
		return res, errAlert
	} else {
		return res, nil
	}
}

func (s *CoinServer) CreateCoins(ctxCon context.Context, in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}

	isExist, errIsExist := db.IsExist(s.db, ctxCon, in.CoinId)
	var err error

	if errIsExist != nil {
		err = fmt.Errorf("There is something wrong with database")
	}
	if isExist {
		err = fmt.Errorf("%v is already exist", in.CoinId)
	} else {
		isCreated := db.InsertOne(s.db, ctxCon, in.CoinId)
		if isCreated {
			res.Status = 201
		} else {
			err = fmt.Errorf("There is no coin named %v", in.CoinId)
		}
	}

	return res, err
}

func (s *CoinServer) SearchCoins(ctxCon context.Context, in *pb.SearchText) (*pb.ReturnList, error) {

	coinList := db.SearchCoins(s.db, ctxCon, in.InputText)

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil
}
