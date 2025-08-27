package apply_code

//
//import (
//	"encoding/json"
//	"net/http"
//	"validator/pkg/logger"
//
//	"github.com/go-chi/chi/v5"
//)
//
//type Handler struct {
//	uc     *ValidatePromoUsecase
//	logger logger.Logger
//}
//
//func NewHandler(uc *ValidatePromoUsecase) *Handler {
//	return &Handler{uc: uc}
//}
//
//// handler.go
//func (h *Handler) ValidateCode(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	code := chi.URLParam(r, "promocode")
//
//	result, err := h.uc.Validate(ctx, code)
//	if err != nil && !result.Exists {
//		http.Error(w, "promocode not found", http.StatusNotFound)
//		return
//	}
//
//	// маппинг домена → DTO (почти 1-в-1)
//	out := Output{
//		Code:       result.Code,
//		Exists:     result.Exists,
//		OnTime:     result.OnTime,
//		Applied:    result.Applied,
//		AppliedNow: result.AppliedNow,
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	_ = json.NewEncoder(w).Encode(out)
//}

//func (h *Handler) ApplyCode(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	var input dto.PromocodeInput
//
//	// map incoming json on dto
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		http.Error(w, "bad request", http.StatusBadRequest)
//		return
//	}
//
//	output_dto, _ := h.uc.ApplyCode(ctx, input)
//
//	out := output_dto
//
//	w.Header().Set("Content-Type", "application/json")
//	_ = json.NewEncoder(w).Encode(out)
//}
