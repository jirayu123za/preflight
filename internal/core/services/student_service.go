package services

import "preflight/internal/core/ports"

type StudentService struct {
	studentRepository ports.StudentRepository
}

var _ ports.StudentService = (*StudentService)(nil)

func NewStudentService(repository ports.StudentRepository) *StudentService {
	return &StudentService{
		studentRepository: repository,
	}
}

func (s *StudentService) Add(studentID int, firstName string, lastName string, email string) error {
	err := s.studentRepository.Add(studentID, firstName, lastName, email)
	if err != nil {
		return err
	}
	return nil
}
