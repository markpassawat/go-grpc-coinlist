package main

import (
	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"
)

func main() {
	db.CreateDatabase()
	db.InsertDefaultCoin()
}
