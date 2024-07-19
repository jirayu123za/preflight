package services

import (
	"preflight/internal/core/repositories"
	"preflight/internal/models"
)

// Primary port
type StudentService interface {
	CreateStudent(student *models.Student) error
	QueryStudentByID(studentID string) (*models.Student, error)
	QueryStudents() ([]models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(student *models.Student) error
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

func (s *StudentServiceImpl) QueryStudentByID(studentID string) (*models.Student, error) {
	// business logic
	student, err := s.repo.FindStudentByStudentId(studentID)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentServiceImpl) QueryStudents() ([]models.Student, error) {
	// business logic
	students, err := s.repo.FindAllStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentServiceImpl) UpdateStudent(student *models.Student) error {
	// business logic
	existingStudent, err := s.repo.FindStudentByStudentId(student.StudentID)
	if err != nil {
		return err
	}

	existingStudent.FirstName = student.FirstName
	existingStudent.LastName = student.LastName
	existingStudent.Email = student.Email

	if err := s.repo.ModifyStudent(existingStudent); err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) DeleteStudent(student *models.Student) error {
	// business logic
	student, err := s.repo.FindStudentByStudentId(student.StudentID)
	if err != nil {
		return err
	}

	if err := s.repo.RemoveStudent(student); err != nil {
		return err
	}
	return nil
}
