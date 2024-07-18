package services

import (
	"preflight/internal/core/repositories"
	"preflight/internal/models"
)

// Primary port
type StudentService interface {
	CreateStudent(student *models.Student) error
	QueryStudents() ([]models.Student, error)
}

type StudentServiceImpl struct {
	repo repositories.StudentRepository
}

// func instance business logic call
func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &StudentServiceImpl{
		repo: repo,
	}
}

func (s *StudentServiceImpl) CreateStudent(student *models.Student) error {
	// business logic
	if err := s.repo.SaveStudent(student); err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) QueryStudents() ([]models.Student, error) {
	// business logic
	students, err := s.repo.FindAllStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}
