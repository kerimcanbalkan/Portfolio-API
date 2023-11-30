package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
	"github.com/kerimcanbalkan/Portfolio-API/middlewares"
)

func MessageRoute(router *gin.Engine) {
	router.POST("/api/messages", controller.CreateMessage)
	// Protected routes with JWT authentication middleware
	adminGroup := router.Group("/api/admin").Use(middlewares.JwtAuthMiddleware())
	{
		adminGroup.GET("/messages", controller.GetMessages)
		adminGroup.GET("/messages/:id", controller.GetMessageById)
		adminGroup.DELETE("/messages/:id", controller.DeleteMessage)
	}
}
