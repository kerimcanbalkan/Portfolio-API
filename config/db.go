package config

import (
	"github.com/kerimcanbalkan/Portfolio-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"))
	if err != nil {
		panic(err)
	}
	DB = db
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Admin{})
}
