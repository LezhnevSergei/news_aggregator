package apiserver

import "github.com/LezhnevSergei/news_aggregator/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
