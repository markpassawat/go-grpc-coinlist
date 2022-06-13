package config

import 	db "github.com/markpassawat/go-grpc-coinlist/pkg/common/db"


type Config struct {
	Environment string `split_words:"true" required:"true"`
	Debug       bool   `split_words:"true" default:"false"`

	Host string `split_words:"true" default:"localhost"`
	Port int    `split_words:"true" default:"8080"`
	DB db.Config `split_words:"true" required:"true"`

}