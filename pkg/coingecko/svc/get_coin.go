package svc

import (
	"context"
	"fmt"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

func (s *CoinServer) GetCoin(ctx context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	res, err := db.GetCoinById(in.CoinId)

	errAlert := fmt.Errorf("There is no coin named %v", in.CoinId)

	if err != nil {
		return res, errAlert
	} else {
		return res, nil
	}
}
