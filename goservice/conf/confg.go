package conf

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
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

	if err := godotenv.Load(ENV_FILE_PATH); err != nil && !os.IsNotExist(err) {
		return Config{}, fmt.Errorf("error leyendo el archivo .env: %w", err)
	}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("could not load environment variables: %w", err)
	}

	return cfg, nil
}
