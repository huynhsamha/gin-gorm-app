package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// CustomBasicModel : customize gorm.Model
type CustomBasicModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

// SetUpDBConnection : used to assign `db` connection
// after connection is established on start server
func SetUpDBConnection(conn *gorm.DB) {
	db = conn
}

// DBTables : Export tables list
// Please order tables to able to delete tables when drop
// Used in migration database
var DBTables = []interface{}{
	&Question{},
	&Answer{},
	&User{},
}
