package service

import (
	"context"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	// Implementation of user registration logic
	newUser = &domain.User{}
	// Save user to repository
	if err := s.repo.Save(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) UpdateProfile(ctx context.Context, id *domain.User) error {
	// Implementation of user profile update logic
	return nil
}

func (s *userService) ChangePassword(ctx context.Context, id *domain.User) error {
	// Implementation of change password logic
	return nil
}

func (s *userService) GetUserProfile(ctx context.Context, id *domain.User) (*domain.User, error) {
	// Implementation of get user profile logic
	return nil, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	// Implementation of user login logic
	return "", nil
}
