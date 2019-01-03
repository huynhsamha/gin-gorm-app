package controllers

import (
	"net/http"

	"github.com/huynhsamha/gin-gorm-app/models"
	"github.com/huynhsamha/gin-gorm-app/utils"
	"github.com/mitchellh/mapstructure"

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
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
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

// Authorized : middleware for user loged in
func (AuthCtrl) Authorized(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	payload, err := jwt.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// store payload (jwtUserData type) to this context
	ctx.Set("payload", payload) // e.g. map[userID:3 username:alice]
	ctx.Next()
}

// getPayload : get `payload` in this context, only used after middleware Authorized()
func (AuthCtrl) getPayload(ctx *gin.Context) (userData jwtUserData, exists bool) {
	payload, ok := ctx.Get("payload") // e.g. map[userID:3 username:alice]
	if !ok {
		return jwtUserData{}, false
	}
	mapPayload := payload.(map[string]interface{})
	result := jwtUserData{}
	mapstructure.Decode(mapPayload, &result) // decode map to struct
	return result, true
}

// CheckAuthorized : check if user loged in, response jwtUserData
func (ctrl AuthCtrl) CheckAuthorized(ctx *gin.Context) {
	payload, ok := ctrl.getPayload(ctx)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request is not authorized"})
		return
	}
	ctx.JSON(http.StatusOK, payload)
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
