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

type Config struct {
	ConfigPath      string
	ServerAddress   string `yaml:"server_address" env:"SERVER_ADDRESS" env-default:":8080"`
	BaseUrl         string `yaml:"base_url" env:"BASE_URL" env-default:"http://localhost:8080"`
	FileStoragePath string `yaml:"file_storage_path" env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `yaml:"database_dsn" env:"DATABASE_DSN"`
}

// TODO: переписать логику обработки конфига
func ParseConfig() (*Config, error) {
	var conf Config

	err := flagParse(&conf)
	if err != nil {
		return nil, err
	}

	envConfPath, ok := os.LookupEnv("CONFIG_FILE")
	if ok {
		conf.ConfigPath = envConfPath
	}

	if _, err := os.Stat(conf.ConfigPath); os.IsNotExist(err) {
		err = cleanenv.ReadEnv(&conf)
		if err != nil {
			return nil, err
		}
		return &conf, ErrConfigFileNotFound
	}

	err = cleanenv.ReadConfig(conf.ConfigPath, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func flagParse(conf *Config) error {

	flag.StringVar(&conf.ConfigPath, "c", "./config/conf.yaml", "path to config file")
	flag.StringVar(&conf.ServerAddress, "a", "", "server address")
	flag.StringVar(&conf.BaseUrl, "b", "", "base url")
	flag.StringVar(&conf.FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&conf.DatabaseDsn, "d", "", "database dsn")

	flag.Parse()
	return nil

}
