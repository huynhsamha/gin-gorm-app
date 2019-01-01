package controllers

import "github.com/jinzhu/gorm"

var db *gorm.DB

// SetUpDBConnection : used to assign `db` connection
// after connection is established on start server
func SetUpDBConnection(conn *gorm.DB) {
	db = conn
}
