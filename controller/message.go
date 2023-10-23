package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	"gorm.io/gorm"
)

func GetMessages(c *gin.Context) {
	messages := []models.Message{}
	config.DB.Find(&messages)
	c.JSON(http.StatusOK, messages)
}

func GetMessageById(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	if err := config.DB.Where("id = ?", id).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, message)
}

func CreateMessage(c *gin.Context) {
	var message models.Message
	c.BindJSON(&message)
	config.DB.Create(&message)
	c.JSON(http.StatusCreated, &message)
}

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	var message models.Message
	result := config.DB.Where("id = ?", id).First(&message)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		}
		return
	}

	if err := config.DB.Delete(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		return
	}

	c.JSON(http.StatusNoContent, "")
}
