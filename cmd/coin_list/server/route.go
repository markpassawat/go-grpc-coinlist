package server

import (
	"context"
	"fmt"

	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	svc "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/svc"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	bun "github.com/uptrace/bun"
)

type CoinServer struct {
	pb.UnimplementedCoinListServer
	db *bun.DB
}

var coins []*pb.CoinInfo

func (s *CoinServer) GetCoins(ctxCon context.Context, in *pb.Empty) (*pb.ReturnList, error) {

	coinList, err := svc.GetAllCoin(s.db, ctxCon)
	return coinList, err

}

func (s *CoinServer) GetCoin(ctxCon context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	coin, err := svc.GetCoinById(s.db, ctxCon, in.CoinId)
	return coin, err

}

func (s *CoinServer) CreateCoins(ctxCon context.Context, in *pb.Id) (*pb.Status, error) {

	err := svc.InsertOne(s.db, ctxCon, in.CoinId)
	res := pb.Status{Status: "99"}
	fmt.Println(&res)
	return &res, err
}

func (s *CoinServer) SearchCoins(ctxCon context.Context, in *pb.SearchText) (*pb.ReturnList, error) {

	coinList := db.SearchCoins(s.db, ctxCon, in.InputText)

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil
}
