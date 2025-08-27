package apply_code

import (
	"encoding/json"
	"net/http"
	"validator/pkg/logger"
)

type Handler struct {
	uc     *UseCase
	logger logger.Logger
}

func NewHandler(uc *UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) ApplyCode(w http.ResponseWriter, r *http.Request) {
	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	result, _ := h.uc.Apply(r.Context(), input)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Output{
		Code:       result.Code,
		Exists:     result.Exists,
		OnTime:     result.OnTime,
		Applied:    result.Applied,
		AppliedNow: result.AppliedNow, // добавляем
	})
}
