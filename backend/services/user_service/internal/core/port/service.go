package port

import (
	"context"
	"user_service/internal/core/domain"
)

type UserService interface {
	// ส่วนที่เปรียบเสมือน COMMAND (Create / Update / Delete)
	Register(ctx context.Context, user *domain.User) error
	UpdateUserProfile(ctx context.Context, req *domain.User) error
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error

	// ส่วนที่เปรียบเสมือน QUERY (Read / View)
	GetUserProfile(ctx context.Context, id uint) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
