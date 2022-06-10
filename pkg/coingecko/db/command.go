package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	pb "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/route"
	coingecko "github.com/superoo7/go-gecko/v3"
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

func IsExist(coinId string) (isExist bool, asd error) {
	db := db.ConnectDatabase()
	ctx := context.TODO()

	isExist, err := db.NewSelect().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Exists(ctx)

	return isExist, err
}

func InsertOne(coinId string) bool {
	ctx := context.TODO()

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

	dbCon := db.ConnectDatabase()

	_, err = dbCon.NewInsert().Model(newCoin).Exec(ctx)

	if err != nil {
		log.Fatal("Err: ", err)
		return false
	} else {
		return true
	}

}

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

func UpdateCoinPrice() {
	ticker := time.NewTicker(time.Hour)
	for _ = range ticker.C {
		db := db.ConnectDatabase()
		ctx := context.TODO()

		// Get id list from database
		dataIdList := new([]Model.Coin)
		idList := []string{}
		errGetId := db.NewSelect().Model((*Model.Coin)(nil)).Order("market_cap_rank ASC").Scan(ctx, dataIdList)

		if errGetId != nil {
			log.Fatal("ERROR: ", errGetId)
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
			log.Fatal("ERROR: ", errGetPrice)
		} else {
			for coinId, coin := range *dataCoinPrice {
				_, errUpdate := db.NewUpdate().Model((*Model.Coin)(nil)).Where("coin_id = ?", coinId).Set("current_price = ?", coin["usd"]).Set("update_at = ?", time.Now()).Exec(ctx)
				if errUpdate != nil {
					log.Fatal("ERROR: ", errUpdate)
				}
			}
			fmt.Println("Coin Update")
		}

		db.Close()
	}
}
