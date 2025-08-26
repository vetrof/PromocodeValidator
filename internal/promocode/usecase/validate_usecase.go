package usecase

import (
	"context"
	"time"
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

func (u *ValidatePromoUsecase) Validate(ctx context.Context, input dto.ValidateInput) (dto.ValidateOutput, error) {
	promo, err := u.repo.GetByCode(ctx, input.Code)
	if err != nil {
		return dto.ValidateOutput{Code: input.Code, Valid: false}, err
	}

	valid := promo.IsValid(time.Now())
	return dto.ValidateOutput{Code: promo.Code, Valid: valid}, nil
}
