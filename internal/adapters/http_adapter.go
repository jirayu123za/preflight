package adapters

import (
	"preflight/internal/core/services"
	"preflight/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Primary adapter
type HttpStudentHandler struct {
	services services.StudentService
}

func NewHttpStudentHandler(services services.StudentService) *HttpStudentHandler {
	return &HttpStudentHandler{
		services: services,
	}
}

func (h *HttpStudentHandler) CreateStudent(c *fiber.Ctx) error {
	// Implement the logic to create a student to the database using GORM.
	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.services.CreateStudent(&student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(student)
}
