package main

import (
	"log"
	"net/http"
	"validator/config"
	"validator/internal/promocode/adapters/fake"
	"validator/internal/promocode/controllers"
	"validator/internal/promocode/usecase"
	"validator/pkg/middleware"
	"validator/pkg/postgres"
	"validator/pkg/token"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
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
	router.Use(chiMiddleware.Logger)

	// auth middleware
	auth := middleware.NewAuth(cfg.JWT.Secret, cfg.JWT.Issuer, cfg.JWT.Audience)

	// token generator + handler
	tokenGen := token.New(cfg.JWT.Secret, cfg.JWT.Issuer, cfg.JWT.Audience)
	tokenHandler := token.NewHandler(tokenGen)

	// юзкейсы и хендлеры
	//promocodeRepo := adapters.NewPgPromoRepo(db)
	promocodeRepo := fake.NewPgPromoRepo(db)
	promocodeUc := usecase.New(promocodeRepo)
	promocodeHandler := controllers.New(promocodeUc)

	// эндпоинт для получения токена
	router.Post("/login", tokenHandler.Login)

	// публичные эндпоинты
	router.Get("/valid_code/{promocode}", promocodeHandler.ValidateCode)

	//// строгая аутентификация (токен обязателен)
	router.Group(func(r chi.Router) {
		r.Use(auth.Middleware(false))
		r.Get("/secure/valid_code/{promocode}", promocodeHandler.ValidateCode)
	})

	log.Println("Server start on :8080 ...")
	http.ListenAndServe(":8080", router)
}
