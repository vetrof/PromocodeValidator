package valid_code

import (
	"encoding/json"
	"net/http"
	"validator/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	uc     *UseCase
	logger logger.Logger
}

func NewHandler(uc *UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) ValidateCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "promocode")

	result, err := h.uc.Validate(ctx, code)
	if err != nil && !result.Exists {
		http.Error(w, "promocode not found", http.StatusNotFound)
		return
	}

	out := Output{
		Code:    result.Code,
		Exists:  result.Exists,
		OnTime:  result.OnTime,
		Applied: result.Applied,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}
