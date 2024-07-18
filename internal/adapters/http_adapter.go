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

func (h *HttpStudentHandler) QueryStudents(c *fiber.Ctx) error {
	// Implement the logic to Query all students to the database using GORM.
	students, err := h.services.QueryStudents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(students)
}

func (h *HttpStudentHandler) QueryStudentByID(c *fiber.Ctx) error {
	// Implement the logic to Query a student by ID to the database using GORM.
	studentID := c.Query("student_id")

	student, err := h.services.QueryStudentByID(studentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(student)
}
