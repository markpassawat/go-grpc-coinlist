package svc

import (
	"context"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/route"
)

func (s *CoinServer) SearchCoins(ctx context.Context, in *pb.SearchText) (*pb.ReturnList, error) {

	coinList := db.SearchCoins(in.InputText)

	coinListReturn := &pb.ReturnList{}
	for _, coin := range coinList {
		coinListReturn.Info = append(coinListReturn.Info, coin)
	}

	return coinListReturn, nil

}
