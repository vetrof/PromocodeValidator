// internal/promocodes/code_validation/validation_handler.go
package http

import (
	"encoding/json"
	"net/http"
	"validator/internal/promocodes/code_validation"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Usecase *code_validation.Usecase
}

func NewHandler(uc *code_validation.Usecase) *Handler {
	return &Handler{Usecase: uc}
}

func (h *Handler) PromocodeValidation(w http.ResponseWriter, r *http.Request) {
	// достаём id из URL /code/{id}
	code := chi.URLParam(r, "id")

	// проверяем код через юзкейс
	result := h.Usecase.ValidCode(code)

	// формируем JSON-ответ
	resp := map[string]any{
		"code":  code,
		"valid": result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
