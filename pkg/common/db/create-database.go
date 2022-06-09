package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	coingecko "github.com/superoo7/go-gecko/v3"
)

func CreateDatabase() {
	dbCon := ConnectDatabase()
	ctx := context.TODO()

	resCreateDB, err := dbCon.NewCreateTable().
		Model((*db.Coin)(nil)).
		Exec(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Create database: ", resCreateDB)

	}
}

func InsertDefaultCoin() {
	db := ConnectDatabase()
	ctx := context.TODO()

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	coinList, err := CG.CoinsMarket("usd", []string{}, "market_cap_desc", 150, 1, false, []string{})

	if err != nil {
		log.Fatal("Err: ", err)
	} else {
		for _, coin := range *coinList {
			coinTemp := &Model.Coin{
				CoinId:        coin.ID,
				Name:          coin.Name,
				Symbol:        coin.Symbol,
				Image:         coin.Image,
				CurrentPrice:  coin.CurrentPrice,
				MarketCapRank: coin.MarketCapRank,
				CreateAt:      time.Now(),
				UpdateAt:      time.Now(),
			}

			_, err := db.NewInsert().Model(coinTemp).Exec(ctx)

			if err != nil {
				log.Fatal("Err: ", err)
			} else {
				fmt.Printf("ADD: %s\n", coin.Name)

			}
		}
	}

}
