package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link-shortner/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
