package logger

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func AddLoggerToGin(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, logger)
		c.Next()
	}
}

func GinLogger(c *gin.Context) {
	start := time.Now()

	// before request
	c.Next()

	// after request
	GetLogger(c).Info("http_request",
		slog.String("method", c.Request.Method),
		slog.String("path", c.Request.URL.Path),
		slog.Int("status", c.Writer.Status()),
		slog.Duration("latency", time.Since(start)),
		slog.String("ip", c.ClientIP()),
	)
}
