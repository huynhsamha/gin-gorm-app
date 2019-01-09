package api

import (
	"github.com/gin-gonic/gin"

	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"
)

var fakeDBCtrl = Controllers.FakeDBCtrl{}

func setUpFakeDBRoutes(router *gin.RouterGroup) {

	authorized := router.Group("/", fakeDBCtrl.AuthRequired)
	authorized.POST("/users", fakeDBCtrl.FakeUsers)
}
