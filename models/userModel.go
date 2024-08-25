package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique"`
	Password string
}
