package db

import (
	"context"

	bun "github.com/uptrace/bun"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
)

func GetAllCoin(db *bun.DB, ctx context.Context) ([]*Model.Coin, error) {

	coinList := new([]*Model.Coin)
	err := db.NewSelect().Model((*Model.Coin)(nil)).Order("market_cap_rank ASC").Scan(ctx, coinList)
	return *coinList, err

}

func GetCoinById(db *bun.DB, ctx context.Context, coinId string) (*Model.Coin, error) {

	coin := new(Model.Coin)
	err := db.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Scan(ctx, coin)
	return coin, err

}
