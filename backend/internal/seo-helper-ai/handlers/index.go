package handlers

import (
	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/templates/views"
	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/utils"
	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return utils.RenderTemplComponent(views.Index("hey"))
}
