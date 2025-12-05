package domain

import "time"

type User struct {
	UserID    int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Address   string
	Role      string
	CreatedAt time.Duration
	UpdatedAt time.Duration
	DeletedAt time.Duration
}
