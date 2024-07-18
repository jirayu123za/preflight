package repositories

import "preflight/internal/models"

// Secondary port
type StudentRepository interface {
	SaveStudent(student *models.Student) error
	//FindStudentByStudentId(studentID int) (*models.Student, error)
	FindAllStudents() ([]models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(student *models.Student) error
}
