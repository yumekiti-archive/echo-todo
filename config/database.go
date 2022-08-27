package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yumekiti/echo-todo/domain"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Task{})

	return db
}
