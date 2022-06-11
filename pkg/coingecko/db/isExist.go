package db

import (
	"context"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
)

func IsExist(coinId string) (isExist bool, asd error) {
	db := db.ConnectDatabase()
	ctx := context.TODO()

	isExist, err := db.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Exists(ctx)

	return isExist, err
}
