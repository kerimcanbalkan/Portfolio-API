package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	"github.com/kerimcanbalkan/Portfolio-API/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.ProjectRoute(router)
	routes.MessageRoute(router)
	router.Run(":8080")
}
