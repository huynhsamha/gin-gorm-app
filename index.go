package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load environment variables in file .env
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var app = gin.Default()

func configApp() {

	/**
	 * Serve static files
	 */
	app.Static("/public", "./public")

	/**
	 * Set up view engine
	 */
	app.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/layout", // master layout for all routes using render HTML
		Partials: []string{ // other partials, allow define to use template
			"partials/index_login",
			"partials/profile_form",
		},
		Funcs: template.FuncMap{
			"nowYear": func() int {
				return time.Now().Year()
			},
		},
		DisableCache: true, // allow reload template on debug without restart server
	})

	/**
	 * View routes
	 */
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

func main() {

	loadEnv()

	configApp()

	port := os.Getenv("port")
	if port == "" {
		port = "3000"
	}
	app.Run(":" + port)
}
