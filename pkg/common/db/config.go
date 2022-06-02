package db

import (
	"fmt"
	"sync"

	"github.com/go-pg/pg/v10"
)

type Config struct {
	Host      string `required:"true"`
	Port      int    `required:"true"`
	User      string `required:"true"`
	Password  string `required:"true"`
	Name      string `required:"true"`
	PoolSize  int    `split_words:"true"`
	AppName   string `split_words:"true"`
	SSLEnable bool   `split_words:"true" envconfig:"SSL_ENABLE"`

	once       sync.Once
	connection *pg.DB
}

func (cfg *Config) GetOpts() *pg.Options {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return &pg.Options{
		Addr:            addr,
		Database:        cfg.Name,
		User:            cfg.User,
		Password:        cfg.Password,
		PoolSize:        cfg.PoolSize,
		ApplicationName: cfg.AppName,
	}
}

func (cfg *Config) Connect() *pg.DB {

	cfg.once.Do(func() {
		cfg.connection = pg.Connect(cfg.GetOpts())
	})

	return cfg.connection
}

func (cfg *Config) MustConnect() *pg.DB {
	conn := cfg.Connect()
	if err := TestConnection(conn); err != nil {
		panic(err)
	}
	return conn
}
