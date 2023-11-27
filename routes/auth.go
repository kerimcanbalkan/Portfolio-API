package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
)

func AuthenticationRoute(router *gin.Engine) {
	router.POST("/api/auth", controller.CreateAdmin)
	router.GET("/api/auth", controller.GetAdmin)
}
