package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

// SetUpDBConnection : used to assign `db` connection
// after connection is established on start server
func SetUpDBConnection(conn *gorm.DB) {
	db = conn
}

// DBTables : Export tables list
// Please order tables to able to delete tables when drop
// Used in migration database
var DBTables = []interface{}{
	&User{},
}
