package model

import "time"

type Coin struct {
	CoinId        string    `bun:",pk,notnull"`
	Name          string    `bun:",notnull"`
	Symbol        string    `bun:",notnull"`
	Image         string    `bun:",notnull"`
	CurrentPrice  float64    `bun:",notnull"`
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