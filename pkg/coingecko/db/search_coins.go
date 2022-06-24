package db

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
)

func SearchCoins(searchText string) []*pb.CoinInfo {
	db := db.ConnectDatabase()
	ctx := context.TODO()

	coinListTemp := new([]*Model.Coin)
	coinList := []*pb.CoinInfo{}
	searchTextTemp := fmt.Sprintf("%%%s%%", searchText)

	err := db.NewSelect().Model((*Model.Coin)(nil)).Order("market_cap_rank ASC").WhereOr("coin_id LIKE ?", searchTextTemp).WhereOr("name LIKE ?", searchTextTemp).WhereOr("symbol LIKE ?", searchTextTemp).Scan(ctx, coinListTemp)

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
