package main

import (
	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Link{})
}
