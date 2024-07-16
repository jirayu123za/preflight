package adapters

import (
	"preflight/internal/core/ports"

	"gorm.io/gorm"
)

type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) ports.StudentRepository {
	return &GormStudentRepository{
		db: db,
	}
}

func (r *GormStudentRepository) Add(int, string, string, string) error {
	// Implement the logic to add a student to the database using GORM.
	return nil
}
