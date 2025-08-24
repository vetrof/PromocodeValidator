package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"validator/internal/member/dto"
	"validator/internal/member/usecase"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) UserExist(w http.ResponseWriter, r *http.Request) {

	user_id_str := chi.URLParam(r, "id")

	// Преобразуем строку в int, обрабатывая возможную ошибку
	user_id_int, err := strconv.Atoi(user_id_str)
	if err != nil {
		// Если преобразование не удалось, возвращаем клиенту ошибку Bad Request
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		log.Printf("Failed to convert user ID: %v", err) // Логируем ошибку для отладки
		return
	}

	result := h.uc.MemberExists(user_id_int)
	response := dto.Output{Username: result}

	// Устанавливаем HTTP-заголовок Content-Type, чтобы клиент знал,
	// что получает JSON.
	w.Header().Set("Content-Type", "application/json")

	// Кодируем структуру `response` в JSON и записываем результат
	// напрямую в `http.ResponseWriter`.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Если произошла ошибка при кодировании JSON,
		// возвращаем внутреннюю ошибку сервера.
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Failed to encode JSON response: %v", err)
		return
	}
}
