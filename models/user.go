package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"index:idx_name,unique"`
	PasswordHash []byte
	Salt         []byte
}
