package api

import (
	Controllers "github.com/huynhsamha/gin-gorm-app/controllers"

	"github.com/gin-gonic/gin"
)

var questionCtrl = Controllers.QuestionCtrl{}

func setUpQuestionRoutes(router *gin.RouterGroup) {
	router.GET("/", questionCtrl.FindAll)
	router.GET("/:id", questionCtrl.FindOneByID)

	authorized := router.Group("/x", authCtrl.Authorized)
	authorized.POST("/create", questionCtrl.Create)
	// authorized.POST("/edit/:id", questionCtrl.EditQuestion)
}
