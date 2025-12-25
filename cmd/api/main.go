package main

import (
	"log"

	"github.com/Dima5791/go-auth-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.Health)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
