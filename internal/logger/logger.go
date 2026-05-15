package logger

import (
	"log/slog"
	"github.com/gin-gonic/gin"
)

const key = "STUGO_LOGGER"

func New() *slog.Logger {
	return slog.Default()
}

func GetLogger(c *gin.Context) *slog.Logger {
	return c.MustGet(key).(*slog.Logger)
}
