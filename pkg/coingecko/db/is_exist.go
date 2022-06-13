package db

import (
	"context"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	bun "github.com/uptrace/bun"
)

func IsExist(db *bun.DB, ctx context.Context,coinId string) (isExist bool, asd error) {

	isExist, err := db.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Exists(ctx)

	return isExist, err
}
