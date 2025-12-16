package port

import (
	"context"
	"user_service/internal/core/domain"
)

type UserRepository interface {
	// Write methods (บันทึกลง DB)
	Save(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, id int, updates map[string]interface{}) error
	Delete(ctx context.Context, id int) error

	// Read methods (ดึงจาก DB)
	AllUsers(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id int) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
