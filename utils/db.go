package utils

import (
	"qnsi/bardun/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func PrepareDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db!")
	}
	db.AutoMigrate(&models.Contact{})
	db.Create(&models.Contact{FirstName: "Bob", LastName: "Kelso"})
	return db
}
