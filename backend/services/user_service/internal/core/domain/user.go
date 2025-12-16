package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique:not null;"`
	Password  string
	Phone     string
	Address   string
	Role      string `gorm:"default:customer;"`
}
