package migrations

import (
	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(&models.Link{})
	initializers.DB.AutoMigrate(&models.User{})
}
