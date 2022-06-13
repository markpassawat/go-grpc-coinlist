package db

import (
	"context"
	"fmt"

	Model "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/model"
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
)

func CreateDatabase() {
	dbCon := db.ConnectDatabase()
	ctx := context.TODO()

	_, err := dbCon.NewCreateTable().
		Model((*Model.Coin)(nil)).
		Exec(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Create database successfully!")

	}
}
