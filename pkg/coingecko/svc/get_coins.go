package svc

import (
	"context"

	query "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	bun "github.com/uptrace/bun"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetAllCoin(db *bun.DB, ctx context.Context) (*pb.ReturnList, error) {

	coinListTemp, err := query.GetAllCoin(db, ctx)
	coinList := &pb.ReturnList{}

	for _, coin := range coinListTemp {
		coinList.Info = append(coinList.Info, &pb.CoinInfo{
			CoinId:        coin.CoinId,
			Name:          coin.CoinId,
			Symbol:        coin.Symbol,
			Image:         coin.Image,
			CurrentPrice:  coin.CurrentPrice,
			MarketCapRank: int32(coin.MarketCapRank),
			CreateAt:      timestamppb.New(coin.CreateAt),
			UpdateAt:      timestamppb.New(coin.UpdateAt),
		})
	}

	return coinList, err

}

func GetCoinById(db *bun.DB, ctx context.Context, coinId string) (*pb.CoinInfo, error) {

	coinTemp, err := query.GetCoinById(db, ctx, coinId)
	coin := &pb.CoinInfo{
		CoinId:        coinTemp.CoinId,
		Symbol:        coinTemp.Symbol,
		Name:          coinTemp.Name,
		Image:         coinTemp.Image,
		CurrentPrice:  coinTemp.CurrentPrice,
		MarketCapRank: int32(coinTemp.MarketCapRank),
		CreateAt:      timestamppb.New(coinTemp.CreateAt),
		UpdateAt:      timestamppb.New(coinTemp.UpdateAt),
	}

	return coin, err

}
