package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"clientIP", c.ClientIP(),
		)
		c.Next()
		slog.InfoContext(c.Request.Context(), "HTTP Response",
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"clientIP", c.ClientIP(),
		)
	}
}
