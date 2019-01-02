package main

import (
	"github.com/huynhsamha/gin-gorm-app/config"
)

/**
 * ***** IMPORTANT *****
 * This function will drop all tables and migrate new tables with gorm
 */

func main() {

	config.LoadEnv()
	config.ConnectDatabase()

	// Drop tables and Migrate new tables
	config.AutoMigrationDB()
}
