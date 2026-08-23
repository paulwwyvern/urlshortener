package config

import (
	"errors"
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	ErrConfigFileNotFound = errors.New("config file not found")
)

type Config struct {
	ConfigPath      string
	ServerAddress   string `yaml:"server_address" json:"server_address"  env:"SERVER_ADDRESS" env-default:":8080"`
	BaseURL         string `yaml:"base_url" json:"base_url" env:"BASE_URL" env-default:"http://localhost:8080"`
	FileStoragePath string `yaml:"file_storage_path" json:"file_storage_path" env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `yaml:"database_dsn" json:"database_dsn" env:"DATABASE_DSN"`
	AuditFile       string `yaml:"audit_file" json:"audit_file" env:"AUDIT_FILE"`
	AuditURL        string `yaml:"audit_url" json:"audit_url" env:"AUDIT_URL"`
	EnableHTTPS     bool   `yaml:"enable_https" json:"enable_https" env:"ENABLE_HTTPS"`
	CertPath        string `yaml:"cert_path" json:"cert_path" env:"CERT_PATH"`
	KeyPath         string `yaml:"key_path" json:"key_path" env:"KEY_PATH"`
	TrustedSubnet   string `yaml:"trusted_subnet" json:"trusted_subnet" env:"TRUSTED_SUBNET"`
}

func ParseConfig() (*Config, error) {
	var conf Config

	err := flagParse(&conf)
	if err != nil {
		return nil, err
	}

	envConfPath, ok := os.LookupEnv("CONFIG")
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
	flag.BoolVar(&conf.EnableHTTPS, "s", false, "enable https")
	flag.StringVar(&conf.ConfigPath, "c", "", "path to config file")
	flag.StringVar(&conf.ConfigPath, "config", "./config/conf.json", "path to config file")
	flag.StringVar(&conf.ServerAddress, "a", "", "server address")
	flag.StringVar(&conf.BaseURL, "b", "", "base url")
	flag.StringVar(&conf.FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&conf.DatabaseDsn, "d", "", "database dsn")
	flag.StringVar(&conf.AuditFile, "audit-file", "", "audit file")
	flag.StringVar(&conf.AuditURL, "audit-url", "", "audit url")
	flag.StringVar(&conf.TrustedSubnet, "t", "", "trusted subnet")

	flag.Parse()
	return nil

}
