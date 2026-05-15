package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"stugo/internal/config"
	"stugo/internal/health"
)

type App struct {
	Config *config.Config
	Router *gin.Engine
}

func NewApp() *App {
	cfg := config.Load()
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	health.Register(router)

	return &App{
		Config: cfg,
		Router: router,
	}
}

func (a *App) Start() {
	port := ":" + a.Config.Port
	router := a.Router

	if err := router.Run(port); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
