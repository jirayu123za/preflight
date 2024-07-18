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

func (r *GormStudentRepository) FindStudentByStudentId(studentID string) (*models.Student, error) {
	// Implement the logic to find a student by studentID from the database using GORM.
	var student models.Student
	if result := r.db.First(&student, "student_id = ?", studentID); result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (r *GormStudentRepository) FindAllStudents() ([]models.Student, error) {
	// Implement the logic to Find all students from the database using GORM.
	var students []models.Student
	if result := r.db.Find(&students); result.Error != nil {
		return nil, result.Error
	}
	return students, nil
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
