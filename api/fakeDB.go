package api

import (
	"github.com/gin-gonic/gin"

	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"
)

var fakeDBCtrl = Controllers.FakeDBCtrl{}

func setUpFakeDBRoutes(router *gin.RouterGroup) {

	router.Use(fakeDBCtrl.AuthRequired)
	router.POST("/users", fakeDBCtrl.FakeUsers)
}
