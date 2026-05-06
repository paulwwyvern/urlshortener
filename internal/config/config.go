package config

import (
	"errors"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

var (
	ErrConfigFileNotFound = errors.New("config file not found")
)

func ParseConfigPath() string {
	var configPath string

	flags := flag.NewFlagSet("config", flag.ContinueOnError)
	flags.StringVar(&configPath, "c", "./config/conf.yaml", "path to config file")
	flags.Parse(os.Args[1:])

	envConfigPath := os.Getenv("CONFIG_FILE")
	if envConfigPath != "" {
		configPath = envConfigPath
	}

	return configPath
}

type Config struct {
	ServerAddress   string `yaml:"server_address" env:"SERVER_ADDRESS" env-default:":8080"`
	BaseUrl         string `yaml:"base_url" env:"BASE_URL" env-default:"http://localhost:8080"`
	FileStoragePath string `yaml:"file_storage_path" env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `yaml:"database_dsn" env:"DATABASE_DSN"`
}

func ParseConfig(path string) (*Config, error) {
	var conf Config

	err := flagParse(&conf)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = cleanenv.ReadEnv(&conf)
		if err != nil {
			return nil, err
		}
		return &conf, ErrConfigFileNotFound
	}

	err = cleanenv.ReadConfig(path, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func flagParse(conf *Config) error {

	flag.StringVar(&conf.ServerAddress, "a", "", "server address")
	flag.StringVar(&conf.BaseUrl, "b", "", "base url")
	flag.StringVar(&conf.FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&conf.DatabaseDsn, "d", "", "database dsn")
	flag.String("c", "", "path to config file")

	flag.Parse()
	return nil

}
