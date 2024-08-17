package seohelperai

import (
	"github.com/Fibig/seo-helper-ai/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	r.GET("/", handleRoute(views.Index("hey")))
	r.GET("/chat", handleRoute(views.SEOHelperAIChat()))

	r.Static("/public", "./assets")
	r.StaticFile("robots.txt", "./assets/robots.txt")
}

func handleRoute(c templ.Component) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.Render(ctx.Request.Context(), ctx.Writer)
	}
}
