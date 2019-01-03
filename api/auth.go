package api

import (
	"github.com/gin-gonic/gin"

	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"
)

var authCtrl = Controllers.AuthCtrl{}

func setUpAuthRoutes(router *gin.RouterGroup) {
	router.POST("/signup", authCtrl.SignUp)
	router.POST("/login", authCtrl.Login)

	authorized := router.Group("/x", authCtrl.Authorized)
	authorized.GET("/authorized", authCtrl.CheckAuthorized)
	authorized.GET("/refreshToken", authCtrl.RefreshToken)
}
