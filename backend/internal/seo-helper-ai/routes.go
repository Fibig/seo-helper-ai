package seohelperai

import (
	"github.com/Fibig/seo-helper-ai/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	// WEB
	r.GET("/", handleRouteTemplComponent(views.Index("hey")))
	r.GET("/chat", handleRouteTemplComponent(views.SEOHelperAIChat()))

	r.Static("/public", "./assets")
	r.StaticFile("robots.txt", "./assets/robots.txt")

	// API
	api := r.Group("/api")
	api.GET("/chat", chatWithAI())
}

func handleRouteTemplComponent(c templ.Component) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.Render(ctx.Request.Context(), ctx.Writer)
	}
}
