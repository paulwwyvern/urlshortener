package config

import (
	"flag"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" `

	BaseUrl string `env:"BASE_URL"`
}

var defaultConfig = Config{
	ServerAddress: ":8080",
	BaseUrl:       "http://localhost:8080",
}

func ParseConfig() (*Config, error) {
	conf := defaultConfig

	err := flagParse(&conf)
	if err != nil {
		return nil, err
	}

	err = envParse(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func envParse(conf *Config) error {
	return env.Parse(conf)

}

func flagParse(conf *Config) error {

	flag.StringVar(&conf.ServerAddress, "a", defaultConfig.ServerAddress, "server address")
	flag.StringVar(&conf.BaseUrl, "b", defaultConfig.BaseUrl, "base url")

	flag.Parse()
	return nil

}
