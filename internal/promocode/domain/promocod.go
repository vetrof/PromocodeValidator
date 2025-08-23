package domain

// PromoCode — чистая доменная модель
type PromoCode struct {
	Code      string
	Valid     bool
	Used      bool
	ValidTill int64 // unix timestamp, 0 = no expiration (пример)
}
