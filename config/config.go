// config/config.go
package config

import (
	"os"
	"strconv"
)

type JWTConfig struct {
	Secret   string
	Issuer   string
	Audience string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Postgres PostgresConfig
	JWT      JWTConfig
}

func Load() Config {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return Config{
		Postgres: PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "secret"),
			DBName:   getEnv("DB_NAME", "app"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:   getEnv("JWT_SECRET", "dev-secret-change-me"),
			Issuer:   os.Getenv("JWT_ISS"), // можно оставить пустым
			Audience: os.Getenv("JWT_AUD"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
