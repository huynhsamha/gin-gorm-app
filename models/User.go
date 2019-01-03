package models

import (
	"errors"

	"github.com/huynhsamha/gin-gorm-app/utils"
	"github.com/jinzhu/gorm"
)

// User : Table name is `users`
type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Email    string `gorm:"unique_index"`
	Password string `gorm:"not null" json:"-"` // json: "-", ignored in responses
	Salt     string `gorm:"not null" json:"-"`

	Name     string
	Location string
	Title    string
	AboutMe  string
	Website  string
	Github   string
	Twitter  string
}

var random = utils.Random{}
var crypto = utils.Crypto{}
var jwt = utils.JWT{}

// GenerateSalt : generate new salt for user, used when sign up or change password
func (user *User) GenerateSalt() {
	user.Salt = random.Hex(32)
}

// HashPassword : hash raw password with salt of user
func (user User) HashPassword(rawPassword string) string {
	return crypto.SHA256(rawPassword + user.Salt)
}

// ValidatePassword : validate if password is correctly
func (user User) ValidatePassword(rawPassword string) bool {
	return user.Password == crypto.SHA256(rawPassword+user.Salt)
}

// ChangePassword : change password of user with old and new password
func (user *User) ChangePassword(rawOldPassword, rawNewPassword string) error {
	if !user.ValidatePassword(rawOldPassword) {
		return errors.New("Password is not correct")
	}
	user.GenerateSalt()
	user.Password = user.HashPassword(rawNewPassword)
	return nil
}
