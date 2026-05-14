package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	app *gin.Engine
}

func (a *App) Run() {
	a.app.Run()
}

func NewApp() *App {
	return &App{
		app: gin.Default(),
	}
}
