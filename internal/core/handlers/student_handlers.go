package handlers

import (
	"preflight/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type StudentHandlers struct {
	studentService ports.StudentService
}

var _ ports.StudentHandlers = (*StudentHandlers)(nil)

func NewStudentHandlers(studentService ports.StudentService) *StudentHandlers {
	return &StudentHandlers{
		studentService: studentService,
	}
}

func (h *StudentHandlers) Add(c *fiber.Ctx) error {
	var studentID int
	var firstName string
	var lastName string
	var email string

	err := h.studentService.Add(studentID, firstName, lastName, email)
	if err != nil {
		return err
	}
	return nil
}

func (h *StudentHandlers) Get(c *fiber.Ctx) error {
	return nil
}

func (h *StudentHandlers) GetAll(c *fiber.Ctx) error {
	return nil
}

func (h *StudentHandlers) Update(c *fiber.Ctx) error {
	return nil
}

func (h *StudentHandlers) Delete(c *fiber.Ctx) error {
	return nil
}
