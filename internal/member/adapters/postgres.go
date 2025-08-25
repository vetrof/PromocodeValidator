package adapters

import (
	"database/sql"
	"errors"
)

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) Exist(id int) (string, error) {
	if id == 99 {
		return "Vasya", nil
	}
	return "", errors.New("member not found")
}
