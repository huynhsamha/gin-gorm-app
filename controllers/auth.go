package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhsamha/gin-gorm-app/models"
)

// AuthCtrl : Controller for Authentication routes
type AuthCtrl struct{}

// SignUp :
func (ctrl AuthCtrl) SignUp(ctx *gin.Context) {
	var form models.User
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, form)
}
