package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type MinioConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	DB           DatabaseConfig
	Minio        MinioConfig
	ServerPort   int
	TokenSecret  string
	PasswordSalt string
}

const (
	COMMON_ENV_PREFIX = "common"
	DB_ENV_PREFIX     = "db"
	MINIO_ENV_PREFIX  = "minio"
)

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(COMMON_ENV_PREFIX, &cfg); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	if err := envconfig.Process(DB_ENV_PREFIX, &cfg.DB); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	if err := envconfig.Process(MINIO_ENV_PREFIX, &cfg.Minio); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	return &cfg, nil
}
