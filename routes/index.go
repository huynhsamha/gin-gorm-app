package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetUp : Set up routes to render view HTML
func SetUp(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title":   "Home Page",
			"appname": "Gin App",
			"message": "Hello Gin Go Web Framework",
		})
	})

	app.GET("/profile", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "profile", gin.H{
			"title":    "Profile Page",
			"appname":  "Gin App",
			"message":  "Hello Gin Go Web Framework",
			"username": "Alice",
		})
	})
}
