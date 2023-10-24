package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(false)
}

type Message struct {
	gorm.Model
	ID      uint   `json:"ID" gorm:"primary_key;unique"`
	Name    string `json:"name"`
	Email   string `json:"email" valid:"email"`
	Message string `json:"message"`
}
