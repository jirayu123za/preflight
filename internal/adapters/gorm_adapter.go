package adapters

import (
	"preflight/internal/core/repositories"
	"preflight/internal/models"

	"gorm.io/gorm"
)

// Secondary adapter
type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) repositories.StudentRepository {
	return &GormStudentRepository{
		db: db,
	}
}

func (r *GormStudentRepository) SaveStudent(student *models.Student) error {
	// Implement the logic to add a student to the database using GORM.
	if result := r.db.Create(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

func convertToStudentSlice(students []*models.Student) []models.Student {
	converted := make([]models.Student, len(students))
	for i, student := range students {
		converted[i] = *student
	}
	return converted
}

func (r *GormStudentRepository) FindAllStudents() ([]models.Student, error) {
	// Implement the logic to Find all students from the database using GORM.
	var students []*models.Student
	if result := r.db.Find(&students); result.Error != nil {
		return nil, result.Error
	}
	return convertToStudentSlice(students), nil
}

func (r *GormStudentRepository) FindStudentByStudentId(student *models.Student) (*models.Student, error) {
	// Implement the logic to Find a student by student id from the database using GORM.
	if result := r.db.First(&student, student.StudentID); result.Error != nil {
		return nil, result.Error
	}
	return student, nil
}

func (r *GormStudentRepository) UpdateStudent(student *models.Student) error {
	// Implement the logic to update a student in the database using GORM.
	if result := r.db.Save(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormStudentRepository) DeleteStudent(student *models.Student) error {
	// Implement the logic to delete a student from the database using GORM.
	if result := r.db.Delete(&student); result.Error != nil {
		return result.Error
	}
	return nil
}
