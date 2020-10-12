package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port int `envconfig:"SERVER_PORT" default:"8000" required:"true"`

	DBName     string `envconfig:"DB_NAME" default:"covid_county"`
	DBUser     string `envconfig:"DB_USER" default:"covid_county"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"secret"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"9999"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"` // disable, require, verify-ca or verify-full
	DBTimeout  int    `envconfig:"DB_TIMEOUT" default:"5"`       // seconds

}

func Load() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("", cfg)

	return cfg, err
}
