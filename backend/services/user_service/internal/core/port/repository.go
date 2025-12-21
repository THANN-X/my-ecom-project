package port

import (
	"context"
	"user_service/internal/core/domain"
)

type UserRepository interface {
	// Write methods (บันทึกลง DB)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, id uint, updates map[string]interface{}) error
	Delete(ctx context.Context, id uint) error

	// Read methods (ดึงจาก DB)
	// AllUsers(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
