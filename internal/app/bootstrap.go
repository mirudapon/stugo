package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"stugo/internal/config"
	"stugo/internal/health"
)

func Start(cfg *config.Config) {
	gin.SetMode(cfg.Mode)
	r := gin.Default()

	health.Register(r)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
