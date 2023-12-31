package main

import (

	// Import "docs" for Swagger documentation generation.
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /api
// @schemes http
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.Use(cors.Default())

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	config.Connect()
	routes.ProjectRoute(router)
	routes.MessageRoute(router)
	routes.ProjectImageRoute(router)
	routes.AuthenticationRoute(router)
	router.Run(":8080")
}
