package routes

import (
	"github.com/gin-gonic/gin"
)

func ProjectImageRoute(router *gin.Engine) {
	router.Static("/project-images", "./upload_directory")
}
