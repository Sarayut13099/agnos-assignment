package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/integralist/go-findroot/find"
	"github.com/joho/godotenv"
)

type Config struct {
	GinMode string `env:"GIN_MODE" envDefault:"debug"`

	AppEnvironment string `env:"APP_ENV" envDefault:"local"`
	AppName        string `env:"APP_NAME" envDefault:"service"`
	AppPort        string `env:"PORT" envDefault:"8080"`

	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"postgres-db"`
	PostgresPort     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER " envDefault:"admin"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"adminpass"`
	PostgresDBName   string `env:"POSTGRES_DBNAME" envDefault:"health_app"`
	PostgresSSLMode  string `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	PostgresTimeZone string `env:"POSTGRES_TIMEZONE" envDefault:"Asia/Bangkok"`

	AccessTokenSigningKey string `env:"ACCESS_TOKEN_SIGNING_KEY" envDefault:"bbaab846a7078d3b5d9e53fc77f30a2f"`
	AccessTokenTTL        int64  `env:"ACCESS_TOKEN_TTL" envDefault:"8"`
}

func New() (*Config, error) {
	cfg := &Config{}
	root, err := find.Repo()
	if err != nil {
		log.Printf("config: unable to find project root: %v", err)
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env", root.Path)); err != nil {
		log.Println("config: no .env file, using system environment")
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
