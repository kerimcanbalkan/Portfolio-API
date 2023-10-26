package controller

import (
	"net/http"
	_ "os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
)

// GetProjects godoc
// @Summary Retrieve all projects
// @Description Get a list of all projects
// @Produce json
// @Success 200 {array} models.Project
// @Router /projects [get]
func GetProjects(c *gin.Context) {
	projects := []models.Project{}
	config.DB.Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// @Summary Create a project
// @Description Create and save a new project with an uploaded image in the database
// @Accept mpfd
// @Produce json
// @Param title formData string true "Title of the project"
// @Param description formData string true "Description of the project"
// @Param image formData file true "Image file to upload"
// @Success 201 {object} jsonResponse
// @Failure 400 {object} jsonResponse
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	// Parse form data to retrieve the uploaded image file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	// Generate a unique filename for the image using a UUID
	fileExt := filepath.Ext(file.Filename)
	filename := uuid.New().String() + fileExt

	// Set the path where you want to save the uploaded image
	imagePath := filepath.Join("upload_directory", filename)

	// Save the uploaded image to the specified path
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Now, you can create a new project record in the database
	// and store the imagePath in the 'Image' field of the Project model
	project := models.Project{
		Image:       imagePath,
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
	}

	// Insert the project into the database using GORM
	config.DB.Create(&project)

	c.JSON(http.StatusCreated, &project)
}
