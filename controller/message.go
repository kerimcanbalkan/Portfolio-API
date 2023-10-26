package controller

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	_ "github.com/kerimcanbalkan/Portfolio-API/docs"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	"gorm.io/gorm"
)

type jsonResponse struct {
	Error string `json:"error"`
}

type CreateMessageRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// GetMessages godoc
// @Summary Retrieve all messages
// @Description Get a list of all messages
// @Produce json
// @Success 200 {array} models.Message
// @Router /messages [get]
func GetMessages(c *gin.Context) {
	messages := []models.Message{}
	config.DB.Find(&messages)
	c.JSON(http.StatusOK, messages)
}

// GetMessageById godoc
// @Summary Retrieve a message by ID
// @Description Get a message by its ID
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} models.Message
// @Failure 404 {object} jsonResponse
// @Router /messages/{id} [get]
func GetMessageById(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	if err := config.DB.Where("id = ?", id).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, message)
}

// CreateMessage godoc
// @Summary Create a message
// @Description Create and save a new message in the database
// @Accept json
// @Produce json
// @Param message body CreateMessageRequest true "Message object"
// @Success 201 {object} models.Message
// @Failure 400 {object} jsonResponse
// @Router /messages [post]
func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform validation using govalidator or other validation package.
	if _, err := govalidator.ValidateStruct(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If validation passes, create the message.
	config.DB.Create(&message)

	c.JSON(http.StatusCreated, &message)
}

// DeleteMessage godoc
// @Summary Delete a message by ID
// @Description Delete a message by its ID
// @Produce json
// @Param id path string true "Message ID"
// @Success 204 "No Content"
// @Failure 404 {object} jsonResponse
// @Failure 500 {object} jsonResponse
// @Router /messages/{id} [delete]
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
