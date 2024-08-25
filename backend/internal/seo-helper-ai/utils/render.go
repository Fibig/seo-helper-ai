package utils

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderTemplComponent(c templ.Component) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.Render(ctx.Request.Context(), ctx.Writer)
	}
}
