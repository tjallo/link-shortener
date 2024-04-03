package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tjalle/link-shortner/controllers"
	"github.com/tjalle/link-shortner/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/links/create", controllers.LinkCreate)

	r.GET("/:link", controllers.LinkGet)

	r.Run()
}
