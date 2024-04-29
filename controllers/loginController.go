package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/auth"
)

type LoginRequestBody struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	hasCorrectCredentials := auth.VerifyUser(requestBody.Username, requestBody.Password)

	if hasCorrectCredentials != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	tokenString, err := auth.CreateToken(requestBody.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
