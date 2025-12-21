package entity

import (
	"user_service/internal/core/domain"

	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique:not null;"`
	Password  string `json:"password" gorm:"not null;"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Role      string `json:"role" gorm:"default:customer;"`
}

func (u *UserEntity) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone,
		Address:   u.Address,
		Role:      u.Role,
	}
}

func FromDomain(user *domain.User) *UserEntity {
	return &UserEntity{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		Address:   user.Address,
		Role:      user.Role,
	}
}
