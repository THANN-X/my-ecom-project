package postgres

import (
	"user_service/internal/core/port"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) port.UserRepository {
	return userRepositoryDB{db: db}
}
