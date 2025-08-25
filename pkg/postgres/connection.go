package postgres

import (
	"database/sql"
	"fmt"
	"validator/config"

	_ "github.com/lib/pq"
)

func NewConnection(cfg config.PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	return sql.Open("postgres", dsn)
}
