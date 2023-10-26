package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Image       string    `json:"image" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null;unique"`
	Description string    `json:"description" gorm:"not null"`
}
