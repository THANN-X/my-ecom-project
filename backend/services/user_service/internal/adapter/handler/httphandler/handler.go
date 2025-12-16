package httphandler

import (
	"user_service/internal/core/domain"
	"user_service/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	// Handler fields and methods
	service port.UserService
}

func NewHttpUserHandler(service port.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

func (h *HttpUserHandler) RegisterUser(c *fiber.Ctx) error {
	// Handler logic for user registration
	user := &domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	createrdUser, err := h.service.Register(c.Context(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createrdUser)
}
