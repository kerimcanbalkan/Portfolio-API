package controller

import (
	"net/http"
	_ "os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	_ "github.com/kerimcanbalkan/Portfolio-API/types"
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

func uploadFile(c *gin.Context) string {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
	}

	fileExt := filepath.Ext(file.Filename)
	filename := uuid.New().String() + fileExt

	imagePath := filepath.Join("upload_directory", filename)

	// Save the uploaded image to the specified path
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
	}

	return imagePath
}

// @Summary Create a project
// @Description Create and save a new project with an uploaded image in the database
// @Accept mpfd
// @Produce json
// @Param title formData string true "Title of the project"
// @Param description formData string true "Description of the project"
// @Param image formData file true "Image file to upload"
// @Success 201 {object} types.AppError
// @Failure 400 {object} types.AppError
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	// // Parse form data to retrieve the uploaded image file
	// file, err := c.FormFile("image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
	// 	return
	// }
	//
	// // Generate a unique filename for the image using a UUID
	// fileExt := filepath.Ext(file.Filename)
	// filename := uuid.New().String() + fileExt
	//
	// // Set the path where you want to save the uploaded image
	// imagePath := filepath.Join("upload_directory", filename)
	//
	// // Save the uploaded image to the specified path
	// if err := c.SaveUploadedFile(file, imagePath); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
	// 	return
	// }

	imagePath := uploadFile(c)

	// Now, you can create a new project record in the database
	// and store the imagePath in the 'Image' field of the Project model
	project := models.Project{
		Image:       imagePath,
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
	}

	project.ID = uuid.New()
	// Insert the project into the database using GORM
	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &project)
}

// @Summary Update a project
// @Description Update an existing project in the database
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param title formData string false "Updated title of the project"
// @Param description formData string false "Updated description of the project"
// @Param image formData file false "Image file to upload"
// @Success 200 {object} models.Project
// @Failure 400 {object} types.AppError
// @Failure 404 "Project not found"
// @Router /projects/{id} [patch]
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var existingProject models.Project

	if err := config.DB.Where("id = ?", id).First(&existingProject).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Check if the title and description are provided in the request

	updatedFilePath := uploadFile(c)
	updatedTitle := c.PostForm("title")
	updatedDescription := c.PostForm("description")

	if updatedFilePath != "upload_directory/" {
		existingProject.Image = updatedFilePath
	}
	if updatedTitle != "" {
		existingProject.Title = updatedTitle
	}

	if updatedDescription != "" {
		existingProject.Description = updatedDescription
	}

	// Update the project in the database
	if err := config.DB.Save(&existingProject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, existingProject)
}
