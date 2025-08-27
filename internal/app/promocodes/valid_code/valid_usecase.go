package valid_code

import (
	"context"
	"time"
	"validator/internal/domain"
)

// PromoRepository — интерфейс репозитория объявлен рядом с usecase.
// Контракт диктует сам usecase, а не база.
type PromoRepository interface {
	GetByCode(ctx context.Context, code string) (*domain.PromoCode, error)
}

type UseCase struct {
	repo PromoRepository
}

func NewUseCase(r PromoRepository) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) Validate(ctx context.Context, code string) (domain.ValidationResult, error) {
	promo, err := u.repo.GetByCode(ctx, code)
	if err != nil {
		// если не нашли — Exists=false
		return domain.NewValidationResult(nil, time.Now()), err
	}

	return domain.NewValidationResult(promo, time.Now()), nil
}
