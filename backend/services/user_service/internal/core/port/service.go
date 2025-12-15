package port

import (
	"context"
	"user_service/internal/core/domain"
)

type UserService interface {
	// ส่วนที่เปรียบเสมือน COMMAND (Write / Action)
	Register(ctx context.Context, email, password string) (*domain.User, error)
	UpdateProfile(ctx context.Context, id *domain.User) error
	ChangePassword(ctx context.Context, id *domain.User) error

	// ส่วนที่เปรียบเสมือน QUERY (Read / View)
	GetUserProfile(ctx context.Context, id *domain.User) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
