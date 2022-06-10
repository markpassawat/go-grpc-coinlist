package main

import (
	db "github.com/markpassawat/go-grpc-coinlist/pkg/coingecko/db"
)

func main() {
	db.CreateDatabase()
	db.InsertDefaultCoin()
}
