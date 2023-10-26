package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
)

func MessageRoute(router *gin.Engine) {
	router.GET("/api/messages", controller.GetMessages)
	router.GET("/api/messages/:id", controller.GetMessageById)
	router.POST("/api/messages", controller.CreateMessage)
	router.DELETE("/api/messages/:id", controller.DeleteMessage)
}
