package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"sync"
)

type Config struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	Password string `required:"true"`
	DB       int    `split_words:"true" required:"true"`
	PoolSize int    `split_words:"true"`

	once   sync.Once
	client *redis.Client
}

func (cfg *Config) GetOpts() *redis.Options {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return &redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	}
}

func (cfg *Config) GetClient() *redis.Client {
	cfg.once.Do(func() {
		cfg.client = redis.NewClient(cfg.GetOpts())
	})
	return cfg.client
}

func (cfg *Config) MustGetClient() *redis.Client {
	client := cfg.GetClient()
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
