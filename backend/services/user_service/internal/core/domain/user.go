package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique:not null;"`
	Password  string `json:"password" gorm:"not null;"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Role      string `json:"role" gorm:"default:customer;"`
}
