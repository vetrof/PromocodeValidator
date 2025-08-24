package main

import (
	"log"
	"net/http"

	"validator/internal/promocode/adapters"
	"validator/internal/promocode/controllers"
	"validator/internal/promocode/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	memberPostgres "validator/internal/member/adapters"
	memberHandler "validator/internal/member/controllers"
	memberUc "validator/internal/member/usecase"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	promocodeRepo := adapters.NewPgPromoRepo()
	promocodeUc := usecase.New(promocodeRepo)
	promocodeHandler := controllers.New(promocodeUc)

	memberRepo := memberPostgres.New()
	memberUc := memberUc.New(memberRepo)
	memberHandler := memberHandler.New(memberUc)

	// GET /code/{id}
	router.Get("/valid_code/{promocode}", promocodeHandler.ValidateCode)
	// GET /member_ex/{id}
	router.Get("/member_ex/{id}", memberHandler.UserExist)

	log.Println("Server start...")
	http.ListenAndServe(":8080", router)
}
