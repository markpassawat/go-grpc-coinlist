package svc

import (
	"context"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

type CoinServer struct {
	pb.UnimplementedCoinListServer
}

var coins []*pb.CoinInfo

func (s *CoinServer) GetCoins(ctx context.Context, in *pb.Empty) (*pb.ReturnList, error) {

	coinList := db.GetAllCoin()

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil

}
