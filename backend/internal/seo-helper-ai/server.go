package seohelperai

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	initRoutes(r)

	return r
}
