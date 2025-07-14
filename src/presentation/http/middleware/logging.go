package middleware

import (
	"time"

	"github.com/eyagovbusiness/GSWB.Users/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Continue to next middleware/handler
		c.Next()

		duration := time.Since(start)

		logger.Logger.Info("HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration.String(),
			"client_ip", c.ClientIP(),
		)
	}
}
