package controllers

import (
	"net/http"

	"github.com/huynhsamha/gin-gorm-app/models"
	"github.com/huynhsamha/gin-gorm-app/utils"

	"github.com/gin-gonic/gin"
)

var jwt = utils.JWT{}

// AuthCtrl : Controller for Authentication routes
type AuthCtrl struct{}

type formSignUp struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
}

// SignUp : validate form, check user exist and create new user
func (ctrl AuthCtrl) SignUp(ctx *gin.Context) {
	var form formSignUp
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Name:     form.Name,
	}
	user.GenerateSalt()
	user.Password = user.HashPassword(form.Password)

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email has been used."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sign up successfully",
		"user":    user,
	})
}

type formLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type jwtUserData struct {
	UserID   uint
	Username string
}

// Login : return jwt (token, iat, nbf, exp, user)
func (AuthCtrl) Login(ctx *gin.Context) {
	var form formLogin
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	res := db.Where(&models.User{Username: form.Username}).First(&user)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username is not existen."})
		return
	}

	if !user.ValidatePassword(form.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password is not correct."})
		return
	}

	token, claims, err := jwt.GenerateToken(jwtUserData{
		UserID:   user.ID,
		Username: user.Username,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
		"iat":   claims.IssuedAt,
		"nbf":   claims.NotBefore,
		"exp":   claims.ExpiresAt,
	})
}

// RefreshToken :
func (AuthCtrl) RefreshToken(ctx *gin.Context) {
	token := ctx.PostForm("token")
	payload, err := jwt.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": payload,
	})
}
