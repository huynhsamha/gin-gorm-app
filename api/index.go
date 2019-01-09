package api

import (
	"github.com/gin-gonic/gin"
)

// SetUp : Set up routes for APIs
func SetUp(router *gin.RouterGroup) {

	// Add fake database routes
	// setUpFakeDBRoutes(router.Group("/fakeDB"))

	// Production API routes
	setUpAuthRoutes(router.Group("/auth"))
	setUpUserRoutes(router.Group("/users"))
}
