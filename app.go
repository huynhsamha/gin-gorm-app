package main

import (
	"html/template"
	"time"

	APIs "github.com/huynhsamha/gin-gorm-app/api"
	"github.com/huynhsamha/gin-gorm-app/config"
	Routes "github.com/huynhsamha/gin-gorm-app/routes"
	"github.com/huynhsamha/gin-gorm-app/utils"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

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
	Routes.SetUp(app)

	/**
	 * API routes
	 */
	api := app.Group("/api")
	APIs.SetUp(api)
}

func main() {

	config.LoadEnv()
	config.ConnectDatabase()

	configApp()

	app.Run(":" + utils.DefaultGetEnv("PORT", "3000"))
}
