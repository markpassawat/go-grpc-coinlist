package db

import (
	"context"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	bun "github.com/uptrace/bun"
)

func InsertOne(db *bun.DB, ctx context.Context, newCoin *Model.Coin) error {

	_, err := db.NewInsert().Model(newCoin).Exec(ctx)
	return err

}
