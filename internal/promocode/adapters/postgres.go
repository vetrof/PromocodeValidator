package adapters

import (
	"context"
	"database/sql"
	"errors"
	"validator/internal/promocode/domain"
)

// PgPromoRepo — заглушка для PostgreSQL.
// На деле хранит моканные данные в map, имитируя таблицу promo_codes.
type PgPromoRepo struct {
	db *sql.DB
}

func NewPgPromoRepo(db *sql.DB) *PgPromoRepo {
	return &PgPromoRepo{db: db}
}

func (r *PgPromoRepo) GetByCode(ctx context.Context, code string) (*domain.PromoCode, error) {
	row := r.db.QueryRowContext(ctx, `SELECT code, expires_at, applied_at FROM promo_codes WHERE code=$1`, code)

	var promo domain.PromoCode
	err := row.Scan(&promo.Code, &promo.ExpiresAt, &promo.AppliedAt)
	if err != nil {
		return nil, err
	}
	return &promo, nil
}

func (r *PgPromoRepo) MarkUsed(ctx context.Context, code string) error {

	return errors.New("promo code not found")
}
