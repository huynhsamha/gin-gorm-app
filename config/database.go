package config

import (
	"fmt"

	"github.com/huynhsamha/gin-gorm-app/controllers"
	"github.com/huynhsamha/gin-gorm-app/models"
	"github.com/huynhsamha/gin-gorm-app/utils"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // driver for connection postgres
)

var db *gorm.DB

// ConnectDatabase : connect to database PostgreSQL using gorm
// gorm : GO - ORM
func ConnectDatabase() {

	dbHost := utils.DefaultGetEnv("DB_HOST", "localhost")
	dbPort := utils.DefaultGetEnv("DB_PORT", "5432")
	dbName := utils.DefaultGetEnv("DB_NAME", "")
	dbUser := utils.DefaultGetEnv("DB_USER", "")
	dbPass := utils.DefaultGetEnv("DB_PASS", "")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass,
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect database PostgreSQL successfully")
	}

	// Pass db connection to package controllers and models
	models.SetUpDBConnection(db)
	controllers.SetUpDBConnection(db)

	// Store this db connection for package config
	setUpDBConnection(db)
}

func setUpDBConnection(DB *gorm.DB) {
	db = DB
}

// Get tables list from Models declaration
var tables = models.DBTables

// AutoMigrationDB : migrate your schema, to keep your schema update to date.
// Document at http://gorm.io/docs/migration.html
// Only used in scripts/migrateDB.go
func AutoMigrationDB() {
	// Drop tables
	db.DropTableIfExists(tables...)
	// Migration
	db.AutoMigrate(tables...)

	// Add Foreign Keys
	db.Model(&models.Profile{}).AddForeignKey("user_id", "users(id)", "SET NULL", "SET NULL")
}
