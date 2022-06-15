package svc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	query "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	coingecko "github.com/superoo7/go-gecko/v3"
	bun "github.com/uptrace/bun"
)

func InsertOne(db *bun.DB, ctx context.Context, coinId string) error {

	isExist, errIsExist := query.IsExist(db, ctx, coinId)

	if errIsExist != nil {
		return errIsExist
	} else {
		if isExist {
			return fmt.Errorf("Already have this coin.")

		} else {

			httpClient := &http.Client{
				Timeout: time.Second * 10,
			}
			CG := coingecko.NewClient(httpClient)

			coin, err := CG.CoinsMarket("usd", []string{coinId}, "market_cap_desc", 1, 1, false, []string{})
			newCoin := &Model.Coin{}

			if err != nil {
				return err
			} else if len(*coin) == 0 {
				return fmt.Errorf("Doesn't have this coin")
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

				err = query.InsertOne(db, ctx, newCoin)

				if err != nil {
					return err
				} else {
					return err
				}
			}

		}
	}

}
