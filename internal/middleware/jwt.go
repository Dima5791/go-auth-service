package middleware

import (
	"net/http"
	"strings"

	"github.com/Dima5791/go-auth-service/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JWT(jwtMgr *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		claims, err := jwtMgr.ValidateAccessToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid access token",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
