package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link-shortner/helpers"
	"github.com/tjalle/link-shortner/initializers"
	"github.com/tjalle/link-shortner/models"
)

func LinkCreate(c *gin.Context) {

	shortLink := helpers.GenerateShortURL(7)

	link := models.Link{
		OriginalLink: "https://www.tjallo.nl/",
		ShortLink:    shortLink,
	}

	result := initializers.DB.Create(&link)

	fmt.Printf("result: %v\n", result)

	fullLink := os.Getenv("BASE_URL") + "/" + shortLink

	c.JSON(http.StatusOK, gin.H{
		"message": fullLink,
	})
}
