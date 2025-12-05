package port

import "user_service/internal/core/domain"

type UserService interface {
	GetUserByID() (*domain.User, error)
}
