package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv : Load environment variables in file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
