package main

import (
	"log"
	"net/http"
	"validator/config"
	"validator/pkg/postgres"

	"validator/internal/promocode/adapters"
	"validator/internal/promocode/controllers"
	"validator/internal/promocode/usecase"

	memberPostgres "validator/internal/member/adapters"
	memberHandler "validator/internal/member/controllers"
	memberUc "validator/internal/member/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// config
	cfg := config.Load()

	// create db connection
	db, err := postgres.NewConnection(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	promocodeRepo := adapters.NewPgPromoRepo(db)
	promocodeUc := usecase.New(promocodeRepo)
	promocodeHandler := controllers.New(promocodeUc)

	memberRepo := memberPostgres.New(db)
	memberUc := memberUc.New(memberRepo)
	memberHandler := memberHandler.New(memberUc)

	// GET /code/{id}
	router.Get("/valid_code/{promocode}", promocodeHandler.ValidateCode)
	// GET /member_ex/{id}
	router.Get("/member_ex/{id}", memberHandler.UserExist)

	log.Println("Server start...")
	http.ListenAndServe(":8080", router)
}
