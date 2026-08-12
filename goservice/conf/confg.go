package conf

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName        string `env:"SERVICE_NAME" envDefault:"local"`
	ServiceVersion     string `env:"SERVICE_VERSION,required"`
	ApplicationPort    string `env:"APPLICATION_PORT" envDefault:":8080"`
	FileCounterGrpcAddr string `env:"FILE_COUNTER_GRPC_ADDR" envDefault:"localhost:5145"`
	Environment        string `env:"ENVIRONMENT" envDefault:"LOCAL"`
	DbConnectionString string `env:"DATABASE_CONN_STRING,required"`
}

func GetEnv() (Config, error) {
	var cfg Config

	if err := godotenv.Load("conf/.env.local"); err != nil {
		return Config{}, err
	}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("could not load environment variables: %w", err)
	}

	return cfg, nil
}
