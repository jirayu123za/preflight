package repositories

import "preflight/internal/models"

type StudentRepository interface {
	CreateStudent(student *models.Student) error
	GetStudentByStudentId(student *models.Student) (*models.Student, error)
	GetAllStudents() ([]models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(student *models.Student) error
}
