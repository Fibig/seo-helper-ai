package main

import (
	"fmt"
	"os"

	seohelperai "github.com/Fibig/seo-helper-ai/internal/seo-helper-ai"
	"github.com/joho/godotenv"
)

func main() {
	// loading dotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}

	server := seohelperai.NewServer()

	server.Run(":" + os.Getenv("PORT"))
}
