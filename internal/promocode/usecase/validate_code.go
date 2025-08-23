package usecase

import (
	"context"
	"validator/internal/promocode/domain"
	"validator/internal/promocode/dto"
)

// PromoRepository интерфейс репозитория объявлен рядом с usecase — usecase диктует контракт
type PromoRepository interface {
	GetByCode(ctx context.Context, code string) (*domain.PromoCode, error)
	MarkUsed(ctx context.Context, code string) error
}

// ValidatePromoUsecase UseCase — реализация application logic
type ValidatePromoUsecase struct {
	repo PromoRepository
}

func New(r PromoRepository) *ValidatePromoUsecase {
	return &ValidatePromoUsecase{repo: r}
}

// Validate проверяет, можно ли использовать промокод
func (u *ValidatePromoUsecase) Validate(ctx context.Context, input dto.ValidateInput) dto.ValidateOutput {
	pc, err := u.repo.GetByCode(ctx, input.Code)
	if err != nil {
		return dto.ValidateOutput{
			Valid:   false,
			Success: false,
			Error:   "промокод не найден",
		}
	}

	if pc == nil {
		return dto.ValidateOutput{
			Valid:   false,
			Success: false,
			Error:   "промокод не найден",
		}
	}

	if pc.Used {
		return dto.ValidateOutput{
			Valid:   false,
			Success: false,
			Error:   "промокод уже использован",
		}
	}

	if !pc.Valid {
		return dto.ValidateOutput{
			Valid:   false,
			Success: false,
			Error:   "промокод просрочен",
		}
	}

	if err := u.repo.MarkUsed(ctx, input.Code); err != nil {
		return dto.ValidateOutput{
			Valid:   false,
			Success: false,
			Error:   "не удалось применить промокод",
		}
	}

	return dto.ValidateOutput{
		Valid:   true,
		Success: true,
	}
}
