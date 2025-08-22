package postgres

// Postgres представляет заглушку репозитория для PostgreSQL.
type Postgres struct{}

func New() *Postgres {
	return &Postgres{}
}

// GetPromocode возвращает заглушку промокода.
func (p *Postgres) ValidCode(code string) string {
	return "qwerty12345"
}
