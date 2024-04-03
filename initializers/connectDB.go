package initializers

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_pass := os.Getenv("MYSQL_PASSWORD")
	mysql_host := os.Getenv("MYSQL_HOST")
	mysql_port := os.Getenv("MYSQL_PORT")
	mysql_db_name := os.Getenv("MYSQL_DB_NAME")

	var sb strings.Builder

	// Constructs a string like this => "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	sb.WriteString(mysql_user)
	sb.WriteString(":")
	sb.WriteString(mysql_pass)
	sb.WriteString("@tcp(")
	sb.WriteString(mysql_host)
	sb.WriteString(":")
	sb.WriteString(mysql_port)
	sb.WriteString(")/")
	sb.WriteString(mysql_db_name)
	sb.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	dsn := sb.String()

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to DB. Are you sure the DB exists and have you ran the migration?")
	}
}
