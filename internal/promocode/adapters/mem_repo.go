package adapters

import (
	"context"
	"errors"
	"validator/internal/promocode/domain"
)

// PgPromoRepo — заглушка для PostgreSQL.
// На деле хранит моканные данные в map, имитируя таблицу promo_codes.
type PgPromoRepo struct {
	data map[string]*domain.PromoCode
}

func NewPgPromoRepo() *PgPromoRepo {
	return &PgPromoRepo{
		data: map[string]*domain.PromoCode{
			"HELLO":   {Code: "HELLO", Valid: true, Used: false},
			"USED":    {Code: "USED", Valid: true, Used: true},
			"EXPIRED": {Code: "EXPIRED", Valid: false, Used: false},
			"NEWYEAR": {Code: "NEWYEAR", Valid: true, Used: false},
		},
	}
}

// GetByCode эмулирует SELECT * FROM promo_codes WHERE code = $1.
func (r *PgPromoRepo) GetByCode(ctx context.Context, code string) (*domain.PromoCode, error) {
	if p, ok := r.data[code]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, errors.New("promo code not found")
}

// MarkUsed эмулирует UPDATE promo_codes SET used=true WHERE code = $1.
func (r *PgPromoRepo) MarkUsed(ctx context.Context, code string) error {
	if p, ok := r.data[code]; ok {
		p.Used = true
		return nil
	}
	return errors.New("promo code not found")
}
