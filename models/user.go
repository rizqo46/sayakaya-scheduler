package models

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	PhoneNumber    string
	Birthday       time.Time `gorm:"type:date"`
	VerifiedStatus bool      `gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}
