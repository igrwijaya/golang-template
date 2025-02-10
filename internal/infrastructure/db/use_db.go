package db

import (
	"os"

	"gorm.io/gorm"
)

type AppDb interface {
	UseMysql() *gorm.DB
	UseSqlite() *gorm.DB
	UseDefault() *gorm.DB
}

type appDb struct {
}

func (appDb *appDb) UseMysql() *gorm.DB {
	return MySqlConfig()
}

func (appDb *appDb) UseSqlite() *gorm.DB {
	return SqliteConfig()
}

func (appDb *appDb) UseDefault() *gorm.DB {

	dbDefault := os.Getenv("DB_DEFAULT")

	if dbDefault == "MYSQL" {
		return MySqlConfig()
	}

	if dbDefault == "SQLITE" {
		return SqliteConfig()
	}

	return MySqlConfig()
}

func NewAppDb() AppDb {
	return &appDb{}
}
