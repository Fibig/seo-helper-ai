package seohelperai

import (
	"github.com/Fibig/seo-helper-ai/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		index := templ.Handler(views.Index())
		index.Component.Render(ctx.Request.Context(), ctx.Writer)
	})

	return r // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
