package main

import (
	"log"
	"net/http"
	http2 "validator/internal/controllers/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"validator/internal/adapters/postgres"
	"validator/internal/promocodes/code_validation"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	repo := postgres.New()
	usecase := code_validation.NewUsecase(repo)
	Handler := http2.NewHandler(usecase)

	// GET /code/{id}
	r.Get("/code/{id}", Handler.PromocodeValidation)

	log.Println("Server start...")
	http.ListenAndServe(":8080", r)
}
