package apply_code

import (
	"context"
	"time"
	"validator/internal/domain"
)

// PromoRepository — интерфейс репозитория объявлен рядом с usecase.
// Контракт диктует сам usecase, а не база.
type PromoRepository interface {
	GetByCode(ctx context.Context, code string) (*domain.PromoCode, error)
	Apply(ctx context.Context, promo *domain.PromoCode, now time.Time) (*domain.PromoCode, error)
}

type UseCase struct {
	repo PromoRepository
}

func NewUseCase(r PromoRepository) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) Apply(ctx context.Context, input Input) (domain.ValidationResult, error) {
	now := time.Now()

	// найти промокод
	promo, err := u.repo.GetByCode(ctx, input.Code)
	if err != nil {
		return domain.NewValidationResult(nil, now), nil
	}

	// провалидировать текущее состояние
	result := domain.NewValidationResult(promo, now)

	// если промо валиден и ещё не применён — применяем
	if result.Exists && result.OnTime && !result.Applied {
		appliedPromo, err := u.repo.Apply(ctx, promo, now)
		if err != nil {
			return result, err
		}
		appliedResult := domain.NewValidationResult(appliedPromo, now)
		appliedResult.AppliedNow = true
		return appliedResult, nil
	}

	// иначе просто вернуть результат
	return result, nil
}
