package repository

import (
	"context"
	"errors"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"gorm.io/gorm"
)

// 1. Struct Declaration: Holds the database connection instance (e.g. *sql.DB or *gorm.DB)
type userRepositoryDB struct {
	db *gorm.DB
}

// 2. Constructor: Creates a new instance of userRepositoryDB
func NewUserRepositoryDB(db *gorm.DB) port.UserRepository {
	return &userRepositoryDB{db: db}
}

// 3. Implement interface methods from port.UserRepository
func (r *userRepositoryDB) AllUsers(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}
	result := r.db.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepositoryDB) Save(ctx context.Context, user *domain.User) error {
	if err := r.db.Create(user); err != nil {
		return err.Error
	}

	return nil
}

func (r *userRepositoryDB) Update(ctx context.Context, id int, updates map[string]interface{}) error {
	user := &domain.User{}
	result := r.db.WithContext(ctx).Model(user).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}
	return nil
}

func (r *userRepositoryDB) Delete(ctx context.Context, id int) error {
	user := &domain.User{}
	result := r.db.WithContext(ctx).Delete(user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}
	return nil
}

func (r *userRepositoryDB) FindById(ctx context.Context, id int) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.WithContext(ctx).First(user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepositoryDB) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	result := r.db.WithContext(ctx).First(user, "email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return user, nil
}
