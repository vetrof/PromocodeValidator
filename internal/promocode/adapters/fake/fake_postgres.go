package fake

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
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
	// Заглушка: эмулируем данные "как будто из базы"
	switch code {
	case "bob":
		return &domain.PromoCode{
			Code:      "bob",
			ExpiresAt: time.Now().Add(24 * time.Hour), // завтра истекает
			AppliedAt: nil,
		}, nil
	case "alice":
		applied := time.Now().Add(-2 * time.Hour)
		return &domain.PromoCode{
			Code:      "alice",
			ExpiresAt: time.Now().Add(24 * time.Hour),
			AppliedAt: &applied, // уже применён
		}, nil
	default:
		// если промокод не найден
		return nil, fmt.Errorf("promo code not found")
	}
}

func (r *PgPromoRepo) MarkUsed(ctx context.Context, code string) error {

	return errors.New("promo code not found")
}
