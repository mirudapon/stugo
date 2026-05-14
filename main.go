package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {

	var app = gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	app.Run()
}

