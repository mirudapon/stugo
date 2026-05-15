package app

import (
	"github.com/gin-gonic/gin"
	"stugo/internal/config"
	"stugo/internal/health"
)

func Start(cfg *config.Config) {
	gin.SetMode(cfg.Mode)
	r := gin.Default()

	health.Register(r)

	r.Run(":" + cfg.Port)
}
