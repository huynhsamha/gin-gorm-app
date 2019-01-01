package api

import (
	"github.com/gin-gonic/gin"

	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"
)

var authCtrl = Controllers.AuthCtrl{}

func setUpAuthRoutes(router *gin.RouterGroup) {
	router.POST("/signup", authCtrl.SignUp)
}
