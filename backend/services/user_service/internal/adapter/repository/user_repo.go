package repository

import (
	"context"
	"errors"
	"user_service/internal/adapter/repository/entity"
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

// 3. Method Implementations:
// Each method interacts with the database using GORM and implements the UserRepository interface
func (r *userRepositoryDB) Create(ctx context.Context, user *domain.User) error {
	// Create a new user
	userEntity := entity.FromDomain(user)

	if err := r.db.WithContext(ctx).Create(userEntity); err != nil {
		return err.Error
	}
	// Map the generated ID back to the domain user
	user.ID = userEntity.ID
	user.CreatedAt = userEntity.CreatedAt
	user.UpdatedAt = userEntity.UpdatedAt

	return nil
}

func (r *userRepositoryDB) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	// Update user by ID
	model := &entity.UserEntity{}
	// Use WithContext to pass the context
	result := r.db.WithContext(ctx).Model(model).Where("id = ?", id).Updates(updates)
	// Handle errors
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *userRepositoryDB) Delete(ctx context.Context, id uint) error {
	// Delete user by ID
	user := &entity.UserEntity{}
	// Use WithContext to pass the context
	result := r.db.WithContext(ctx).Delete(user, id)
	// Handle errors
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepositoryDB) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	// Find user by ID
	user := &entity.UserEntity{}
	// Use WithContext to pass the context
	result := r.db.WithContext(ctx).First(user, id)
	// Handle errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return user.ToDomain(), nil
}

func (r *userRepositoryDB) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	// Find user by email
	user := &entity.UserEntity{}
	// Use WithContext to pass the context
	result := r.db.WithContext(ctx).First(user, "email = ?", email)
	// Handle errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return user.ToDomain(), nil
}
