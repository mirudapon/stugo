package app

import (
	"github.com/gin-gonic/gin"
	"stugo/interval/health"
)

func Start() {
	var app = gin.Default()

	health.Register(app)

	app.Run()
}
