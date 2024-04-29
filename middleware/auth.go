package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/auth"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Request.Header["Token"]

		if !ok || len(token) == 0 || token[0] == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No JWT token found"})
			c.Abort()
			return
		}

		isValidToken := auth.VerifyToken(token[0])

		if isValidToken != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token sent"})
			c.Abort()
			return
		}

		c.Next()
	}
}
