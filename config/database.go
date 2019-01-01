package config

import (
	"fmt"

	"github.com/huynhsamha/gin-gorm-app/utils"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // driver for connection postgres
)

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
	defer db.Close()

}
