package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/helpers"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Token"][0]

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No JWT token found"})
			c.Abort()
			return
		}

		isValidToken := helpers.VerifyToken(token)

		if isValidToken != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token sent"})
			c.Abort()
			return
		}

		c.Next()
	}
}
