package conf

import (
	"errors"
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName         string `env:"SERVICE_NAME" envDefault:"local"`
	ServiceVersion      string `env:"SERVICE_VERSION,required"`
	ApplicationPort     string `env:"APPLICATION_PORT" envDefault:":8080"`
	FileCounterGrpcAddr string `env:"FILE_COUNTER_GRPC_ADDR" envDefault:"localhost:5145"`
	Environment         string `env:"ENVIRONMENT" envDefault:"LOCAL"`
	DbConnectionString  string `env:"DATABASE_CONN_STRING,required"`
}

const ENV_FILE_PATH = "conf/.env"

func GetEnv() (Config, error) {
	var cfg Config

	_, err := os.Stat(ENV_FILE_PATH)

	if errors.Is(err, os.ErrNotExist) {
		cfg = Config{
			ServiceName:         os.Getenv("SERVICE_NAME"),
			ServiceVersion:      os.Getenv("SERVICE_VERSION"),
			ApplicationPort:     os.Getenv("APPLICATION_PORT"),
			FileCounterGrpcAddr: os.Getenv("FILE_COUNTER_GRPC_ADDR"),
			Environment:         os.Getenv("ENVIRONMENT"),
			DbConnectionString:  os.Getenv("DATABASE_CONN_STRING"),
		}

		return cfg, err
	}

	if err := godotenv.Load(ENV_FILE_PATH); err != nil {
		return Config{}, err
	}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("could not load environment variables: %w", err)
	}

	return cfg, nil
}
