package db

import (
	"context"
	"fmt"
	"net/http"
	"time"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	coingecko "github.com/superoo7/go-gecko/v3"
	bun "github.com/uptrace/bun"
)

func UpdateCoinPrice(db *bun.DB, ctx context.Context) {
	ticker := time.NewTicker(time.Hour)
	for range ticker.C {

		// Get id list from database
		dataIdList := new([]Model.Coin)
		idList := []string{}
		errGetId := db.NewSelect().Model((*Model.Coin)(nil)).Order("market_cap_rank ASC").Scan(ctx, dataIdList)

		if errGetId != nil {
			fmt.Println("ERROR: ", errGetId)
		} else {
			for _, coin := range *dataIdList {
				idList = append(idList, coin.CoinId)
			}
		}

		httpClient := &http.Client{
			Timeout: time.Second * 10,
		}
		CG := coingecko.NewClient(httpClient)

		dataCoinPrice, errGetPrice := CG.SimplePrice(idList, []string{"usd"})

		if errGetPrice != nil {
			fmt.Println("ERROR: ", errGetPrice)
		} else {
			for coinId, coin := range *dataCoinPrice {
				_, errUpdate := db.NewUpdate().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Set("current_price = ?", coin["usd"]).Set("update_at = ?", time.Now()).Exec(ctx)
				if errUpdate != nil {
					fmt.Println("ERROR: ", errUpdate)
				}
			}
			fmt.Println("Coin Update")
		}

		db.Close()
	}
}
