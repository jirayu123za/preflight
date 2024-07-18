package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID string `json:"student_id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
}
