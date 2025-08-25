package main

import (
	"log"
	"net/http"
	"validator/config"
	"validator/pkg/middleware"
	"validator/pkg/postgres"
	"validator/pkg/token"

	"validator/internal/promocode/adapters"
	"validator/internal/promocode/controllers"
	"validator/internal/promocode/usecase"

	memberPostgres "validator/internal/member/adapters"
	memberHandler "validator/internal/member/controllers"
	memberUc "validator/internal/member/usecase"

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
	promocodeRepo := adapters.NewPgPromoRepo(db)
	promocodeUc := usecase.New(promocodeRepo)
	promocodeHandler := controllers.New(promocodeUc)

	memberRepo := memberPostgres.New(db)
	memberUc := memberUc.New(memberRepo)
	memberHandler := memberHandler.New(memberUc)

	// эндпоинт для получения токена
	router.Post("/login", tokenHandler.Login)

	// публичные эндпоинты
	router.Get("/valid_code/{promocode}", promocodeHandler.ValidateCode)
	router.Get("/member_ex/{id}", memberHandler.UserExist)

	// защищённые эндпоинты
	router.Group(func(r chi.Router) {
		r.Use(auth.Middleware)
		r.Get("/strong_auth/", memberHandler.StrongValidationToken)
	})

	log.Println("Server start on :8080 ...")
	http.ListenAndServe(":8080", router)
}
