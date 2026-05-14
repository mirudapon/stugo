package app

import (
	"github.com/gin-gonic/gin"
	"stugo/interval/health"
)

func Start(port string, mode string) {
	gin.SetMode(mode)
	var app = gin.Default()

	health.Register(app)

	app.Run(":" + port)
}
