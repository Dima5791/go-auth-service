package server

import (
	"github.com/Dima5791/go-auth-service/internal/handler"
	"github.com/Dima5791/go-auth-service/internal/middleware"
	"github.com/Dima5791/go-auth-service/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	authHandler *handler.AuthHandler,
	jwtMgr *jwt.Manager,
) *gin.Engine {

	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)
	r.POST("/logout", authHandler.Logout)

	profileHandler := handler.NewProfileHandler()

	auth := r.Group("/")
	auth.Use(middleware.JWT(jwtMgr))
	{
		auth.GET("/profile", profileHandler.Profile)
	}

	adminHandler := handler.NewAdminHandler()

	admin := r.Group("/admin")
	admin.Use(
		middleware.JWT(jwtMgr),
		middleware.RequireRole("admin"),
	)
	{
		admin.GET("/dashboard", adminHandler.Dashboard)
	}

	return r
}

func RegisterRoutes(
	r *gin.Engine,
	authHandler *handler.AuthHandler,
) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}
}
