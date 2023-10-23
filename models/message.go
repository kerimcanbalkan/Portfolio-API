package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id      int    `json:"ID" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
