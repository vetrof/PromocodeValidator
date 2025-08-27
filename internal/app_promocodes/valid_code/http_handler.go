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
	code := chi.URLParam(r, "promocode")

	result, err := h.uc.Validate(r.Context(), code)

	w.Header().Set("Content-Type", "application/json")

	// If there's an error and the code doesn't exist, use a default output.
	if err != nil && !result.Exists {
		_ = json.NewEncoder(w).Encode(Output{
			Code:   code,
			Exists: false,
		})
		return
	}

	// Otherwise, use the result from the use case.
	_ = json.NewEncoder(w).Encode(Output{
		Code:    result.Code,
		Exists:  result.Exists,
		OnTime:  result.OnTime,
		Applied: result.Applied,
	})
}
