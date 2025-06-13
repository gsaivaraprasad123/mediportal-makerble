package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string         `gorm:"not null"`
	Age       int            `gorm:"not null"`
	Gender    string         `gorm:"not null"`
	Condition string         `gorm:"not null"`
	UpdatedBy string         `gorm:"type:varchar(100)"` // Refers to doctor who last updated
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
