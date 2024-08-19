package seohelperai

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()

	initRoutes(r)

	return r
}
