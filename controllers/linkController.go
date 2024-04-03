package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link-shortner/helpers"
	"github.com/tjalle/link-shortner/initializers"
	"github.com/tjalle/link-shortner/models"
)

func LinkCreate(c *gin.Context) {
	shortLink := helpers.GenerateShortURL(7)

	// TODO: Add way to check if post param is actually an URL
	originalUrl := c.PostForm("url")
	isEmptyPostForm := originalUrl == ""

	if isEmptyPostForm {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing 'url' parameter",
		})

		return
	}

	link := models.Link{
		OriginalLink: originalUrl,
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

	fullLink := os.Getenv("BASE_URL") + "/" + shortLink

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": fullLink,
	})
}
