package token

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	gen *Generator
}

func NewHandler(gen *Generator) *Handler {
	return &Handler{gen: gen}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// ⚡ Тут должна быть реальная проверка юзера (например, через БД)
	if req.Login != "admin" || req.Password != "secret" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// ⚡ Допустим, у юзера id=123 и роль admin
	tok, err := h.gen.Generate("123", []string{"admin"}, time.Hour)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tok})
}
