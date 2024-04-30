package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/auth"
	"github.com/tjalle/link_shortener/controllers"
	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/middleware"
	"github.com/tjalle/link_shortener/migrations"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
	migrations.Migrate()
}

func main() {
	r := gin.New()

	auth.CreateUser("tjalle", "SuperLangWachtwoord12!")

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
