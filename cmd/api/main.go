package main

import (
	"database/sql"
	"log"

	"github.com/Dima5791/go-auth-service/internal/config"
	"github.com/Dima5791/go-auth-service/internal/handler"
	"github.com/Dima5791/go-auth-service/internal/repository"
	"github.com/Dima5791/go-auth-service/internal/server"
	"github.com/Dima5791/go-auth-service/internal/service"
	"github.com/Dima5791/go-auth-service/pkg/jwt"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	refreshRepo := repository.NewRefreshTokenRepo(db)

	jwtMgr := jwt.NewManager(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessTTL,
		cfg.JWT.RefreshTTL,
	)

	authService := service.NewAuthService(
		userRepo,
		refreshRepo,
		jwtMgr,
	)

	authHandler := handler.NewAuthHandler(authService)

	r := server.NewRouter(authHandler, jwtMgr)
	r.Run(":8080")
}
