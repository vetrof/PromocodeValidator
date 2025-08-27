package fake

import (
	"context"
	"errors"
	"fmt"
	"time"
	"validator/internal/domain"
)

type PgPromoRepo struct {
	data map[string]*domain.PromoCode
}

func NewFakePostgres() *PgPromoRepo {
	return &PgPromoRepo{
		data: map[string]*domain.PromoCode{
			"bob": {
				Code:      "bob",
				ExpiresAt: time.Now().Add(24 * time.Hour),
				AppliedAt: nil,
			},
			"alice": {
				Code:      "alice",
				ExpiresAt: time.Now().Add(24 * time.Hour),
				AppliedAt: func() *time.Time {
					t := time.Now().Add(-2 * time.Hour)
					return &t
				}(),
			},
		},
	}
}

func (r *PgPromoRepo) GetByCode(ctx context.Context, code string) (*domain.PromoCode, error) {
	promo, ok := r.data[code]
	if !ok {
		return nil, fmt.Errorf("promo code not found")
	}
	return promo, nil
}

func (r *PgPromoRepo) Apply(ctx context.Context, promo *domain.PromoCode, now time.Time) (*domain.PromoCode, error) {
	if promo == nil {
		return nil, errors.New("nil promo")
	}
	promo.Apply(now)           // меняем объект в памяти
	r.data[promo.Code] = promo // обновляем "хранилище"
	return promo, nil
}
