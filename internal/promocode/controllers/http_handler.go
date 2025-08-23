package controllers

import (
	"encoding/json"
	"net/http"

	"validator/internal/promocode/dto"
	"validator/internal/promocode/usecase"
	"validator/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	uc     *usecase.ValidatePromoUsecase
	logger logger.Logger
}

func New(uc *usecase.ValidatePromoUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) ValidateCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	promocode := chi.URLParam(r, "promocode")

	output := h.uc.Validate(ctx, dto.ValidateInput{Code: promocode})

	if !output.Success {
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = json.NewEncoder(w).Encode(output)
}

//func (h *Handler) ApplyCode(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	var req dto.ValidateInput
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		h.logger.With("handler", "ApplyCode").Error("bad_request", "err", err)
//		http.Error(w, "bad request", http.StatusBadRequest)
//		return
//	}
//
//	h.logger.With("handler", "ApplyCode", "code", req.Code).Info("request")
//	ok, err := h.uc.Apply(ctx, req.Code)
//	if err != nil {
//		h.logger.With("handler", "ApplyCode").Error("internal_error", "err", err)
//		http.Error(w, "internal error", http.StatusInternalServerError)
//		return
//	}
//	_ = json.NewEncoder(w).Encode(map[string]any{"applied": ok})
//}
