package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null;unique"`
}

func (a *Admin) BeforeCreate() error {
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)

	return nil
}
