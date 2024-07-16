package adapters

import (
	"preflight/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type HttpStudentHandler struct {
	services *services.StudentService
}

func NewHttpStudentHandler(services *services.StudentService) *HttpStudentHandler {
	return &HttpStudentHandler{
		services: services,
	}
}

func (h *HttpStudentHandler) Add(c *fiber.Ctx) error {
	// Implement the logic to add a student to the database using GORM.
	return nil
}
