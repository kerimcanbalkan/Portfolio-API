package controller

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	_ "github.com/kerimcanbalkan/Portfolio-API/docs"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	_ "github.com/kerimcanbalkan/Portfolio-API/types"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

// GetMessages godoc
// @Summary Retrieve all messages
// @Description Get a list of all messages
// @Produce json
// @Success 200 {array} models.Message
// @Router /messages [get]
// @Tags Messages
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
// @Failure 404 {object} types.AppError
// @Router /messages/{id} [get]
// @Tags Messages
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
// @Param message body types.CreateMessageRequest true "Message object"
// @Success 201 {object} models.Message
// @Failure 400 {object} types.AppError
// @Router /messages [post]
// @Tags Messages
func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.ID = uuid.New()

	// Perform validation using govalidator or other validation package.
	if _, err := govalidator.ValidateStruct(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	// Send an email with the message content
	if err := sendEmail(message); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &message)
}

func sendEmail(message models.Message) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "kerimcanbalkanportfolio@gmail.com")
	m.SetHeader("To", "kerimcanbalkan@gmail.com")
	m.SetHeader("Subject", message.Name+"/ "+message.Email)
	m.SetBody("text/plain", "A new message has been created:\n\n"+message.Message)

	d := gomail.NewDialer("smtp.gmail.com", 587, "kerimcanbalkanportfolio@gmail.com", "jwfw mcnm oxvi mcrp")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// DeleteMessage godoc
// @Summary Delete a message by ID
// @Description Delete a message by its ID
// @Produce json
// @Param id path string true "Message ID"
// @Success 204 "No Content"
// @Failure 404 {object} types.AppError
// @Failure 500 {object} types.AppError
// @Router /messages/{id} [delete]
// @Tags Messages
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
