package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/ping", func (c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	http.ListenAndServe(":8080", router)
}
