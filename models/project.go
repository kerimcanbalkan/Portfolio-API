package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Id          uint   `json:"ID" gorm:"primary_key;unique"`
	Image       string `json:"image" gorm:"not null"`
	Title       string `json:"title" gorm:"not null;unique"`
	Description string `json:"description" gorm:"not null"`
}
