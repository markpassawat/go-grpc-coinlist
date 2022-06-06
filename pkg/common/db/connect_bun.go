package db

import (
	"context"
	"fmt"

	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/model"
	"github.com/uptrace/bun"
)

func CreateTableDB(dbCon *bun.DB, ctxCon context.Context) {
	res_createDBCoin, err := dbCon.NewCreateTable().Model((*db.Coin)(nil)).Exec(ctxCon)
	if err != nil {
		fmt.Println("Cannot create: ", err)
	} else {
		fmt.Println("res_createDBCoin => ", res_createDBCoin)
	}
}

func CreateDatabase() {
	dbCon := ConnectDatabase()
	ctx := context.TODO()

	_, err := dbCon.NewCreateTable().
		Model((*db.Coin)(nil)).
		Exec(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect database successfully!")

	}
}
