package httphandler

import (
	"strconv"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type ChangePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserHandler struct {
	// Handler fields and methods
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Handler for getting user profile
func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	// Handler logic for getting user profile
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	// Call service to get user profile
	user, err := h.svc.GetUserProfile(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Handler for user registration
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	// Handler logic for user registration
	user := &domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Call service to register user
	err := h.svc.Register(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// Handler for user login
func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	// Handler logic for user login
	user := &domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Call service to login user
	msg, err := h.svc.Login(c.Context(), user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to login user",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(msg)
}

// Handler for changing password
func (h *UserHandler) ChangePassword(c *fiber.Ctx) error {
	// Handler logic for changing password
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	// Call service to change password
	req := &ChangePasswordReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Call service to change password
	err = h.svc.ChangePassword(c.Context(), uint(id), req.OldPassword, req.NewPassword)
	if err != nil {
		if err.Error() == "incorrect old password" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect old password"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to change password",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

// Handler for updating user profile
func (h *UserHandler) UpdateUserProfile(c *fiber.Ctx) error {
	// Handler logic for updating user profile
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	// Call service to update user profile
	user := &domain.User{}
	// Parse request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Set user ID from URL parameter
	user.ID = uint(id)
	// Call service to update user profile
	err = h.svc.UpdateUserProfile(c.Context(), user)
	// Handle errors
	if err != nil {
		// Specific error handling for "user not found"
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		// General error handling
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}
