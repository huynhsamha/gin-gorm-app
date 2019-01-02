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
// gorm (GO ORM for SQL): http://gorm.io/docs/connecting_to_the_database.html
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

// GetDBConnection : get db connection from package config
func GetDBConnection() *gorm.DB {
	return db
}
