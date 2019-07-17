package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB = dbInit()

type User struct {
	gorm.Model
	Email          string
	PasswordSalted string
}

func dbInit() *gorm.DB {
	// Setup a database.
	// TODO(jaykhatri) figure out how to limit the number of open
	// db connections.
	db, err := gorm.Open("sqlite3", "test.db")
	db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
