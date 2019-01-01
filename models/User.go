package models

import (
	"github.com/jinzhu/gorm"
)

// User : Table name is `users`
type User struct {
	gorm.Model
	Username string `gorm:"unique_index" form:"username" json:"username" binding:"required"`
	Email    string `gorm:"unique_index" form:"email" json:"email" binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required"`
	Salt     string `gorm:"not null"`
}

// Profile : Table name is `profiles`
// Each profile belongs to one user
type Profile struct {
	gorm.Model
	UserID   int
	User     User // belongs to User
	Name     string
	Location string
	Title    string
	AboutMe  string
	Website  string
	Github   string
	Twitter  string
}
