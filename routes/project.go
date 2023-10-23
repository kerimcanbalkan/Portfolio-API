package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/controller"
)

func ProjectRoute(router *gin.Engine) {
	router.GET("/projects", controller.ProjectController)
}
