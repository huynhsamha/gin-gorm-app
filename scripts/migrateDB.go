package main

import (
	"fmt"

	"github.com/huynhsamha/gin-gorm-app/config"
	"github.com/huynhsamha/gin-gorm-app/models"
)

/**
 * ***** IMPORTANT *****
 * This function will drop all tables and migrate new tables with gorm
 */

// Get tables list from Models declaration
var tables = models.DBTables

func main() {

	config.LoadEnv()
	config.ConnectDatabase()

	// Drop tables and Migrate new tables
	autoMigrationDB()
}

// autoMigrationDB : migrate your schema, to keep your schema update to date.
// Document at http://gorm.io/docs/migration.html
// Only used in scripts/migrateDB.go
func autoMigrationDB() {

	// Get DB connnection
	db := config.GetDBConnection()

	// Drop tables
	fmt.Println("Dropping tables...")
	db.DropTableIfExists(tables...)

	// Migration
	fmt.Println("Migrating database...")
	db.AutoMigrate(tables...)

	// Add Foreign Keys
	fmt.Println("Add foreign keys...")
	// db.Model(&models.Profile{}).AddForeignKey("user_id", "users(id)", "SET NULL", "SET NULL")

	fmt.Println("Migration is successfully")
}
