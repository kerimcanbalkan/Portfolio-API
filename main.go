package main

import (

	// Import "docs" for Swagger documentation generation.
	"github.com/gin-gonic/gin"
	"github.com/kerimcanbalkan/Portfolio-API/config"
	_ "github.com/kerimcanbalkan/Portfolio-API/docs"
	"github.com/kerimcanbalkan/Portfolio-API/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Portfolio API
// @version 1.0
// @description Api for my portfolio website

// @host localhost:8080
// @BasePath /api
func main() {
	router := gin.New()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	config.Connect()
	routes.ProjectRoute(router)
	routes.MessageRoute(router)
	router.Run(":8080")
}
