package ports

import "github.com/gofiber/fiber/v2"

type StudentService interface {
	Add(studentID int, firstName string, lastName string, email string) error
}

type StudentRepository interface {
	Add(studentID int, firstName string, lastName string, email string) error
}

type StudentHandlers interface {
	Add(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
