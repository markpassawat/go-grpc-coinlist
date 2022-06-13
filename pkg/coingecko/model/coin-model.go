package model

import (
	"time"
	PB "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	bun "github.com/uptrace/bun"
)

type Coin struct {
	CoinId        string    `bun:",pk,notnull"`
	Name          string    `bun:",notnull"`
	Symbol        string    `bun:",notnull"`
	Image         string    `bun:",notnull"`
	CurrentPrice  float64   `bun:",notnull"`
	MarketCapRank int16     `bun:",notnull"`
	CreateAt      time.Time `bun:",nullzero,default:now()"`
	UpdateAt      time.Time `bun:",nullzero,default:now()"`
}

type UpdateCoin struct {
	CoinId        string    `bun:",pk,notnull"`
	CurrentPrice  string    `bun:",notnull"`
	MarketCapRank string    `bun:",notnull"`
	UpdateAt      time.Time `bun:",nullzero,default:now()"`
}

type CoinServer struct {
	PB.UnimplementedCoinListServer
	DB *bun.DB
}