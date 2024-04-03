package main

import (
	"github.com/tjalle/link-shortner/initializers"
	"github.com/tjalle/link-shortner/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Link{})
}
