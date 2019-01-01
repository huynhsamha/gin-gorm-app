package models

import (
	"github.com/jinzhu/gorm"
)

// User : Table name is `users`
type User struct {
	gorm.Model
	username string `gorm:"unique_index"`
	email    string `gorm:"unique_index"`
	password string `gorm:"not null"`
	salt     string `gorm:"not null"`
}
