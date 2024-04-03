package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/helpers"
	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/models"
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

	tx := initializers.DB.Begin()

	if err := tx.Create(&link).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create shortened URL, please try again.",
		})
		return
	}

	tx.Commit()

	fullURL := os.Getenv("BASE_URL") + "/" + shortLink

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": fullURL,
	})
}

func LinkGet(c *gin.Context) {
	link := c.Param("link")

	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
		})
		return
	}

	var l models.Link

	result := initializers.DB.Where("short_link = ?", link).First(&l)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "link not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, l.OriginalLink)
}
