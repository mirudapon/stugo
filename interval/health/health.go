package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Register(r *gin.Engine) {
	r.GET("/ping", ping)
}
