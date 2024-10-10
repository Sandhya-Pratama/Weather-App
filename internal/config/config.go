package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

//Config adalah konfigurasi yang digunakan
type Config struct {
	Port     string         `env:"PORT" envDefault:"8080"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
}

//PostgresConfig merupakan konfigurasi ke postgres
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}

func NewConfig(envPath string) (*Config, error) {
	// Memuat file environment dari envPath menggunakan godotenv
	if err := godotenv.Load(envPath); err != nil {
		return nil, errors.New("failed to load environment file: " + err.Error())
	}

	// Parsing variabel environment ke dalam struct Config
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("failed to parse environment variables: " + err.Error())
	}

	return &cfg, nil
}