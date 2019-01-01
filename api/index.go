package api

import (
	"github.com/gin-gonic/gin"
)

// SetUp : Set up routes for APIs
func SetUp(router *gin.RouterGroup) {
	setUpAuthRoutes(router.Group("/auth"))
	setUpUserRoutes(router.Group("/users"))
}
