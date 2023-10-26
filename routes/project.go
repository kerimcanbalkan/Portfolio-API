package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
)

func ProjectRoute(router *gin.Engine) {
	router.GET("/api/projects", controller.GetProjects)
	router.POST("/api/projects", controller.CreateProject)
	router.Static("/project-images", "./upload_directory")
	router.Run(":8081")
}
