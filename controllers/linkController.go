package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link-shortner/helpers"
	"github.com/tjalle/link-shortner/initializers"
	"github.com/tjalle/link-shortner/models"
)

func LinkCreate(c *gin.Context) {
	shortLink := helpers.GenerateShortURL(7)

	originalURL := c.PostForm("url")
	if originalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing 'url' parameter",
		})
		return
	}

	// Validate the URL
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
		})
		return
	}

	link := models.Link{
		OriginalLink: originalURL,
		ShortLink:    shortLink,
	}

	result := initializers.DB.Create(&link)
	// Poor man's duplicate checking, could (and maybe should) be improved
	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "There was an error creating this URL, please try again.",
			})
		return
	}

	fullURL := os.Getenv("BASE_URL") + "/" + shortLink

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": fullURL,
	})
}
