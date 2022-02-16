package models

import (
	"time"

	"gorm.io/gorm"
)

func GetContacts(db *gorm.DB) []Contact {
	var contacts []Contact
	db.Find(&contacts)
	return contacts
}

type Contact struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
}
