package db

import (
	"context"
	"log"
	"net/http"
	"time"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	bun "github.com/uptrace/bun"
	coingecko "github.com/superoo7/go-gecko/v3"
)

func InsertOne(db *bun.DB, ctx context.Context,coinId string) bool {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	coin, err := CG.CoinsMarket("usd", []string{coinId}, "market_cap_desc", 1, 1, false, []string{})
	newCoin := &Model.Coin{}

	if err != nil {
		log.Fatal("ERR:", err)
		return false
	} else if len(*coin) == 0 {
		return false
	} else {
		newCoin = &Model.Coin{
			CoinId:        coinId,
			Name:          (*coin)[0].Name,
			Symbol:        (*coin)[0].Symbol,
			Image:         (*coin)[0].Image,
			CurrentPrice:  (*coin)[0].CurrentPrice,
			MarketCapRank: (*coin)[0].MarketCapRank,
			CreateAt:      time.Now(),
			UpdateAt:      time.Now(),
		}
	}


	_, err = db.NewInsert().Model(newCoin).Exec(ctx)

	if err != nil {
		log.Fatal("Err: ", err)
		return false
	} else {
		return true
	}

}
