package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
)

func MessageRoute(router *gin.Engine) {
	router.GET("/messages", controller.GetMessages)
	router.GET("/messages/:id", controller.GetMessageById)
	router.POST("/messages", controller.CreateMessage)
	router.DELETE("/messages/:id", controller.DeleteMessage)
}
