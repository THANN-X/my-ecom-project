package domain

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Address   string
	Role      string
}
