package svc

import (
	"context"
	"fmt"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/route"
)


func (s *CoinServer) CreateCoins(ctx context.Context, in *pb.Id) (*pb.Status, error) {
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
