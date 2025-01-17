package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config adalah konfigurasi yang digunakan
type Config struct {
	Env      string         `env:"ENV" envDefault:"development"`
	Port     string         `env:"PORT" envDefault:"8080"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	JWT      JwtConfig      `envPrefix:"JWT_"`
	Session  SessionConfig  `envPrefix:"SESSION_"`
	Redis    RedisConfig    `envPrefix:"REDIS_"`
}

//	type MidtransConfig struct {
//		BaseURL   string `env:"BASE_URL"`
//		ServerKey string `env:"SERVER_KEY"`
//		ClientKey string `env:"CLIENT_KEY"`
//	}

type RedisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"6379"`
	Password string `env:"PASSWORD" envDefault:""`
}

// JWT CONFIG
type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}
type SessionConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

// PostgresConfig merupakan konfigurasi ke postgres
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}

// NewConfig creates a new config
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// parseConfig parses the configuration file located at envPath and returns a
// Config struct and an error if any. It uses the godotenv package to load the
// environment variables from the file and the env package to parse them into
// the Config struct.
//
// envPath: The path to the environment file.
// Returns: A pointer to the Config struct and an error.
func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}
