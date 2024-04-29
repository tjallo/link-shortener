package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/controllers"
	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/middleware"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authorized := r.Group("/")

	authorized.Use(middleware.JWTAuthMiddleware())

	// Authorized Routes
	{
		authorized.POST("/links/create", controllers.LinkCreate)

		authorized.GET("/getAll", controllers.LinkGetAll)
	}

	// Unauthorized Routes
	{
		r.POST("/login", controllers.Login)

		r.GET("/:link", controllers.LinkGet)
	}

	r.Run()
}
