package api

import (
	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"

	"github.com/gin-gonic/gin"
)

var userCtrl = Controllers.UserCtrl{}

func setUpUserRoutes(router *gin.RouterGroup) {
	router.GET("/", userCtrl.FindAll)
	router.GET("/n/:username", userCtrl.FindOneByUsername)
	router.GET("/d/:id", userCtrl.FindOneByID)
}
