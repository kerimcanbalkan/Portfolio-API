package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
)

func CreateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the BeforeCreate hook to hash the password
	if err := admin.BeforeCreate(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Save the admin to the database
	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin user"})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

func GetAdmin(c *gin.Context) {
	admin := []models.Admin{}
	config.DB.Find(&admin)
	c.JSON(http.StatusOK, admin)
}

func Login(c *gin.Context) {
}
