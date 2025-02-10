package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteConfig() *gorm.DB {
	db, openDbError := gorm.Open(sqlite.Open("template.db"), &gorm.Config{})

	if openDbError != nil {
		panic("Can't connect to database")
	}

	return db
}
