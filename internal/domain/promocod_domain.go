package domain

import "time"

type PromoCode struct {
	Code         string
	ExpiresAt    time.Time
	AppliedAt    *time.Time // nil, если не применён
	MerchantName string
	Phone        string
}

type ValidationResult struct {
	Code       string
	Exists     bool
	OnTime     bool
	Applied    bool
	AppliedNow bool
}

// Уже применён?
func (p *PromoCode) IsApplied() bool {
	return p.AppliedAt != nil
}

// Действителен по времени?
func (p *PromoCode) IsOnTime(now time.Time) bool {
	return now.Before(p.ExpiresAt)
}

func NewValidationResult(promo *PromoCode, now time.Time) ValidationResult {
	if promo == nil {
		return ValidationResult{Exists: false}
	}
	return ValidationResult{
		Code:       promo.Code,
		Exists:     true,
		OnTime:     promo.IsOnTime(now),
		Applied:    promo.IsApplied(),
		AppliedNow: promo.AppliedNow(now),
	}
}

func (p *PromoCode) AppliedNow(now time.Time) bool {
	return p.IsApplied() && p.AppliedAt != nil && p.AppliedAt.After(now.Add(-5*time.Minute))
}
