package config

type Config struct {
	Environment string `split_words:"true" required:"true"`
	Debug       bool   `split_words:"true" default:"false"`

	Host string `split_words:"true" default:"localhost"`
	Port int    `split_words:"true" default:"8080"`

	AllowedOrigin string `split_words:"true" default:"*"`

	Moviesapp string `split_words:"true" default:"localhost:8081"`
}
