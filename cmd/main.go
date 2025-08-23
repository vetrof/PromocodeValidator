package main

import (
	"log"
	"net/http"

	"validator/internal/promocode/adapters"
	"validator/internal/promocode/controllers"
	"validator/internal/promocode/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	promocodeRepo := adapters.NewPgPromoRepo()
	promocodeUc := usecase.New(promocodeRepo)
	promocodeHandler := controllers.New(promocodeUc)

	// GET /code/{id}
	router.Get("/valid_code/{promocode}", promocodeHandler.ValidateCode)

	log.Println("Server start...")
	http.ListenAndServe(":8080", router)
}
