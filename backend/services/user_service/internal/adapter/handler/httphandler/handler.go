package httphandler

import (
	"strconv"
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// Handler fields and methods
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := h.svc.GetUserProfile(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	// Handler logic for user registration
	user := &domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := h.svc.Register(c.Context(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
