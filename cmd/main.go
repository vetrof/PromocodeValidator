// todo Дописать применение промокода

package main

import (
	"log"
	"net/http"
	"validator/config"
	"validator/internal/adapters/postgres/fake"
	apply_code2 "validator/internal/app_promocodes/apply_code"
	valid_code2 "validator/internal/app_promocodes/valid_code"
	"validator/internal/controllers"
	"validator/pkg/token"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// config
	cfg := config.Load()

	// create db connection
	//db, err := postgres.NewConnection(cfg.Postgres)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// create repo
	postgresRepo := fake.NewFakePostgres()

	// main router
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)

	// token generator
	tokenGen := token.New(cfg.JWT.Secret, cfg.JWT.Issuer, cfg.JWT.Audience)
	tokenHandler := token.NewHandler(tokenGen)
	router.Post("/login", tokenHandler.Login)

	// APP /promocode/ //
	// /valid code
	validCodeUseCase := valid_code2.NewUseCase(postgresRepo)
	validCodeHandler := valid_code2.NewHandler(validCodeUseCase)
	// /apply code
	applyCodeUseCase := apply_code2.NewUseCase(postgresRepo)     // todo
	applyCodeHandler := apply_code2.NewHandler(applyCodeUseCase) // todo
	// router
	controllers.PromocodesRouter(router, validCodeHandler, applyCodeHandler) // todo

	log.Println("Server start on :8080 ...")
	http.ListenAndServe(":8080", router)
}
