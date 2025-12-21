package service

import (
	"context"
	"errors"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, user *domain.User) error {
	// Implementation of user registration logic
	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user to repository
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserProfile(ctx context.Context, req *domain.User) error {
	// Implementation of user profile update logic
	// user, err := s.repo.FindByID(ctx, req.ID)
	// if err != nil {
	// 	return err
	// }

	// Update user fields
	updates := map[string]interface{}{}
	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if len(updates) == 0 {
		return nil // No updates to make
	}
	// Save updated user to repository
	if err := s.repo.Update(ctx, req.ID, updates); err != nil {
		return err
	}
	return nil
}

func (s *userService) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	// Implementation of change password logic
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("incorrect old password") // Old password does not match
	}
	// Validate new password (example: minimum length)
	if len(newPassword) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Update password in repository
	changepass := map[string]interface{}{
		"password": string(hashedPassword),
	}
	// Save updated password to repository
	return s.repo.Update(ctx, userID, changepass)
}

func (s *userService) GetUserProfile(ctx context.Context, id uint) (*domain.User, error) {
	// Implementation of get user profile logic
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	// Implementation of user login logic
	// Find user by email
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return "Login successful", nil
}
