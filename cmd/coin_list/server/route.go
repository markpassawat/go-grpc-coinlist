package server

import (
	"context"
	"fmt"

	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
	Model "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	pb "github.com/markpassawat/go-grpc-coinlist/proto/coinlist"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type coinServer struct {
	pb.UnimplementedCoinListServer
}

var coins []*pb.CoinInfo

func (s *coinServer) GetCoins(in *pb.Empty, stream pb.CoinList_GetCoinsServer) error {

	coinlist := db.GetAll()

	for _, coin := range coinlist {
		if err := stream.Send(coin); err != nil {
			return err
		}
	}
	return nil

}

func (s *coinServer) GetCoin(ctx context.Context,
	in *pb.Id) (*pb.CoinInfo, error) {

	res := &pb.CoinInfo{}

	// Connect DB
	dbCon := db.ConnectDatabase()

	// GET
	coinTemp := new(Model.Coin)
	errGetCoinByID := dbCon.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", in.CoinId).Scan(ctx, coinTemp)

	if errGetCoinByID != nil {
		fmt.Println("Don't have coin id: ", in.CoinId)
	} else {
		res.CoinId = coinTemp.CoinId
		res.Symbol = coinTemp.Symbol
		res.Name = coinTemp.Name
		res.Image = coinTemp.Image
		res.CurrentPrice = coinTemp.CurrentPrice
		res.MarketCapRank = int32(coinTemp.MarketCapRank)
		res.CreateAt = timestamppb.New(coinTemp.CreateAt)
		res.UpdateAt = timestamppb.New(coinTemp.UpdateAt)
	}

	return res, nil
}

func (s *coinServer) CreateCoins(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}

	dbCon := db.ConnectDatabase()

	// POST
	exists, err := dbCon.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", in.CoinId).Exists(ctx)
	if err != nil {
		panic(err)
	}
	if exists {
		res.Status = 409
	} else {
		db.InsertOne(in.CoinId)
		res.Status = 201

	}

	return res, nil
}

func (s *coinServer) UpdateCoins(ctx context.Context,
	in *pb.CoinInfo) (*pb.Status, error) {

	res := &pb.Status{}
	for index, coin := range coins {
		if coin.GetCoinId() == in.GetCoinId() {
			coins = append(coins[:index], coins[index+1:]...)
			in.CoinId = coin.GetCoinId()
			coins = append(coins, in)
			res.Status = 200
			break
		}
	}

	return res, nil
}

func (s *coinServer) DeleteCoin(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	res := &pb.Status{}
	res.Status = 204
	for index, coin := range coins {
		if coin.GetSymbol() == in.GetSymbol() {
			coins = append(coins[:index], coins[index+1:]...)
			res.Status = 200
			break
		}
	}

	return res, nil
}

func (s *coinServer) SearchCoins(in *pb.InputText, stream pb.CoinList_SearchCoinsServer) error {

	coinlist := db.SearchCoins(in.InputText)

	for _, coin := range coinlist {
		if err := stream.Send(coin); err != nil {
			return err
		}
	}
	return nil
}
