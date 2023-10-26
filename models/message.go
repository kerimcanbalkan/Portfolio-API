package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(false)
}

type Message struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	Name    string    `json:"name"`
	Email   string    `json:"email" valid:"email"`
	Message string    `json:"message"`
}
