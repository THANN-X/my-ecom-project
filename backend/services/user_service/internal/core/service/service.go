package service

import (
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) port.UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]domain.User, error) {
	return nil, nil
}

func (r userRepositoryDB) GetById() (*domain.User, error) {
	return nil, nil
}
