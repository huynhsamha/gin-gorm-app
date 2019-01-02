package controllers

import (
	"net/http"

	"github.com/huynhsamha/gin-gorm-app/models"
	"github.com/huynhsamha/gin-gorm-app/utils"

	"github.com/gin-gonic/gin"
)

var random = utils.Random{}
var crypto = utils.Crypto{}

// AuthCtrl : Controller for Authentication routes
type AuthCtrl struct{}

type formSignUp struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// SignUp :
func (ctrl AuthCtrl) SignUp(ctx *gin.Context) {
	var form formSignUp
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
		Salt:     random.Hex(32),
	}
	user.Password = crypto.SHA256(user.Password + user.Salt)

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sign up successfully",
		"user":    user,
	})
}
