package main

import (
	"fmt"
	"os"

	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/handlers"
	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// loading dotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}

	// setting gin mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// server init
	server := gin.New()

	// WEB
	server.GET("/", handlers.Index())
	server.GET("/generate", handlers.Generate())

	// API
	api := server.Group("/api")
	api.POST("/generate", handlers.APIGenerate())

	server.Use(middlewares.CacheMiddleware(os.Getenv("GIN_MODE") == "debug"))
	server.Static("/public", "./public")
	server.StaticFile("robots.txt", "./public/robots.txt")

	server.Run(":" + os.Getenv("PORT"))
}
