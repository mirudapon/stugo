package app

import (
	"github.com/gin-gonic/gin"
	"stugo/internal/config"
	"stugo/internal/health"
)

func Start(cfg *config.Config) {
	gin.SetMode(cfg.Mode)
	var app = gin.Default()

	health.Register(app)

	app.Run(":" + cfg.Port)
}
