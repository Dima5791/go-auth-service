package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct{}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}

func (h *ProfileHandler) Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
	})
}
