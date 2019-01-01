package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDatabase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=gorm password=mypassword")
	defer db.Close()
}
