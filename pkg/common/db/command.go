package db

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"


	// "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	Model "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	coingecko "github.com/superoo7/go-gecko/v3"
)

func InsertOne(coinId string) {
	// db := ConnectDatabase()
	ctx := context.TODO()

	// log := logrus.StandardLogger()

	// Create and Update 95 coins in database  ------------->
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	// coinListDetails := []*Model.Coin{}
	coin, err := CG.CoinsMarket("usd", []string{coinId}, "market_cap_desc", 1, 1, false, []string{})
	newCoin := &Model.Coin{}

	if err != nil {
		panic(err)
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
		// return newCoin
		// for _, coinTemp := range *coin {
		// 	newCoin =  Model.Coin{
		// 		CoinId:        coinId,
		// 		Name:          coin.Name,
		// 		Symbol:        coin.Symbol,
		// 		Image:         coin.Image,
		// 		CurrentPrice:  fmt.Sprint(coin.CurrentPrice),
		// 		MarketCapRank: coin.MarketCapRank,
		// 		CreateAt:      time.Now(),
		// 		UpdateAt:      time.Now(),
		// 	}
		// }
	}

	dbCon := ConnectDatabase()

	_, err = dbCon.NewInsert().Model(newCoin).Exec(ctx)

	if err != nil {
		log.Fatal("Err: ", err)
	}

}

func GetAll() []*pb.CoinInfo {
	db := ConnectDatabase()
	ctx := context.TODO()
	log := logrus.StandardLogger()

	coinListTemp := new([]*Model.Coin)
	coinList := []*pb.CoinInfo{}

	err := db.NewSelect().Model((*Model.Coin)(nil)).Order("market_cap_rank ASC").Scan(ctx, coinListTemp)

	for _, coinTemp := range *coinListTemp {
		coinList = append(coinList, &pb.CoinInfo{
			CoinId:        coinTemp.CoinId,
			Name:          coinTemp.CoinId,
			Symbol:        coinTemp.Symbol,
			Image:         coinTemp.Image,
			CurrentPrice:  coinTemp.CurrentPrice,
			MarketCapRank: int32(coinTemp.MarketCapRank),
			CreateAt:      timestamppb.New(coinTemp.CreateAt),
			UpdateAt:      timestamppb.New(coinTemp.UpdateAt),
		})
	}

	if err != nil {
		log.Fatal("Err: ", err)
	} else {
		return coinList
	}

	return nil

}
