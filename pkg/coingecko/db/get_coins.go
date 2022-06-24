package db

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

func GetAllCoin() []*pb.CoinInfo {
	db := db.ConnectDatabase()
	ctx := context.TODO()

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

func GetCoinById(coinId string) (*pb.CoinInfo, error) {
	db := db.ConnectDatabase()
	ctx := context.TODO()

	coinTemp := new(Model.Coin)

	errGetCoinByID := db.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Scan(ctx, coinTemp)

	coin := &pb.CoinInfo{}

	if errGetCoinByID == nil {
		coin.CoinId = coinTemp.CoinId
		coin.Symbol = coinTemp.Symbol
		coin.Name = coinTemp.Name
		coin.Image = coinTemp.Image
		coin.CurrentPrice = coinTemp.CurrentPrice
		coin.MarketCapRank = int32(coinTemp.MarketCapRank)
		coin.CreateAt = timestamppb.New(coinTemp.CreateAt)
		coin.UpdateAt = timestamppb.New(coinTemp.UpdateAt)
		return coin, errGetCoinByID
	}

	return coin, errGetCoinByID

}
