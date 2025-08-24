package adapters

import "errors"

type Postgres struct{}

func New() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Exist(id int) (string, error) {
	if id == 99 {
		return "Vasya", nil
	}
	return "", errors.New("member not found")
}
