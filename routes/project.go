package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
	"github.com/kerimcanbalkan/Portfolio-API/middlewares"
)

func ProjectRoute(router *gin.Engine) {
	router.GET("/api/projects", controller.GetProjects)
	adminGroup := router.Group("/api/admin").Use(middlewares.JwtAuthMiddleware())
	{
		adminGroup.POST("/projects", controller.CreateProject)
		adminGroup.PATCH("/projects", controller.UpdateProject)
	}
}
