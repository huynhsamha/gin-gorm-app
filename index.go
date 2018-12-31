package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app = gin.Default()

func configApp() {

	// Serve static files
	app.Static("/", "./public")

}

func main() {

	// Load environment variables in file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configApp()

	port := os.Getenv("port")
	if port == "" {
		port = "3000"
	}
	app.Run(":" + port)
}
