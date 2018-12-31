package main

import (
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func configApp() {

	// Serve static files
	app.Static("/", "./public")
}

func main() {
	configApp()

	app.Run(":8080")
}
