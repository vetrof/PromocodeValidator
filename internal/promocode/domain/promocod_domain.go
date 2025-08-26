package domain

import "time"

type PromoCode struct {
	Code      string
	ExpiresAt time.Time
	AppliedAt *time.Time // nil если не применён
}

// IsValid Можно использовать?
func (p *PromoCode) IsValid(now time.Time) bool {
	if now.After(p.ExpiresAt) {
		return false
	}
	if p.AppliedAt != nil {
		return false
	}
	return true
}
