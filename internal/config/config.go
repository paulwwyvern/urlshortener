package config

import (
	"flag"
)

type Config struct {
	ServerAddress string

	UrlShortenerAddress string
}

func ParseConfig() *Config {
	conf := &Config{}
	flag.StringVar(&conf.ServerAddress, "a", ":8080", "server address")
	flag.StringVar(&conf.UrlShortenerAddress, "b", "http://localhost:8080", "url shortener address")

	flag.Parse()
	return conf
}
