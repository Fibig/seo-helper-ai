package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CacheMiddleware(devMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := "public, max-age=604800, immutable"

		if devMode {
			headerValue = "no-cache"
		}

		c.Writer.Header().Set("Cache-Control", headerValue)

		c.Next()
	}
}
