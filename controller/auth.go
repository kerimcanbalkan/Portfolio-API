package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/models"
	"github.com/kerimcanbalkan/Portfolio-API/types"
	_ "github.com/kerimcanbalkan/Portfolio-API/types"
	"github.com/kerimcanbalkan/Portfolio-API/utils"
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

// Login godoc
// @Summary Authenticate admin user
// @Description Authenticate admin user and generate JWT token
// @Accept json
// @Produce json
// @Param admin body types.LoginInput true "Admin credentials"
// @Success 200 {object} types.LoginResponse "JWT token"
// @Failure 404 {object} types.AppError
// @Router /auth [post]
// @Tags Admin
func Login(c *gin.Context) {
	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.LoginCheck(admin.Username, admin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong username or password"})
		return
	}

	c.JSON(http.StatusOK, types.LoginResponse{Token: token})
}
