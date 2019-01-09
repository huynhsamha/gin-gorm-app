package controllers

import (
	"net/http"

	"github.com/huynhsamha/gin-gorm-app/utils"

	"github.com/huynhsamha/gin-gorm-app/models"

	"github.com/gin-gonic/gin"
)

// FakeDBCtrl : Controller for fake database routes
type FakeDBCtrl struct{}

// AuthRequired : require authorization to fake database if server allows fake APIs
func (FakeDBCtrl) AuthRequired(ctx *gin.Context) {
	keyFakeDB, exist := utils.RequireGetEnv("KEY_FAKE_DB")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "API is not supported"})
		return
	}

	if ctx.GetHeader("Authorizarion") != keyFakeDB {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Key authorization is not correct"})
		return
	}

	ctx.Next()
}

type formFakeUser struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`

	Location string `form:"location" json:"location"`
	Title    string `form:"title" json:"title"`
	AboutMe  string `form:"aboutMe" json:"aboutMe"`
	Website  string `form:"website" json:"website"`
	Github   string `form:"github" json:"github"`
	Twitter  string `form:"twitter" json:"twitter"`
	PhotoURL string `form:"photoUrl" json:"photoUrl"`
}

// FakeUsers : fake database users
func (ctrl FakeDBCtrl) FakeUsers(ctx *gin.Context) {
	var form formFakeUser
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: form.Username,
		Email:    form.Email,
		Name:     form.Name,
		Location: form.Location,
		Title:    form.Title,
		AboutMe:  form.AboutMe,
		Website:  form.Website,
		Github:   form.Github,
		Twitter:  form.Twitter,
		PhotoURL: form.PhotoURL,
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
