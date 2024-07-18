package repositories

import "preflight/internal/models"

// Secondary port
type StudentRepository interface {
	SaveStudent(student *models.Student) error
	FindStudentByStudentId(studentID string) (*models.Student, error)
	FindAllStudents() ([]models.Student, error)
	ModifyStudent(student *models.Student) error
	DeleteStudent(student *models.Student) error
}
