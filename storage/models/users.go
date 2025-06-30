package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"not null"; unique`
	Username       string
	HashedPassword string
}

type Store struct {
	gorm.Model
	UserID    uint
	StoreName string
}
