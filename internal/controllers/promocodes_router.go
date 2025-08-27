package controllers

import (
	"validator/internal/app_promocodes/apply_code"
	"validator/internal/app_promocodes/valid_code"

	"github.com/go-chi/chi/v5"
)

// PromocodesRouter отвечает за настройку маршрутов, связанных с промокодами.
func PromocodesRouter(r chi.Router,
	validHandler *valid_code.Handler,
	applyHandler *apply_code.Handler,
) {
	// Вся логика роутинга для промокодов здесь
	r.Route("/promocode", func(r chi.Router) {
		// Публичные эндпоинты
		r.Get("/validate/{promocode}", validHandler.ValidateCode)
		r.Post("/apply/", applyHandler.ApplyCode)
	})
}
