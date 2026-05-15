package app

import (
	"log/slog"
	"os"
	"github.com/gin-gonic/gin"
	"stugo/internal/config"
	"stugo/internal/health"
	"stugo/internal/logger"
)

type App struct {
	Config *config.Config
	Router *gin.Engine
	Logger *slog.Logger
}

func NewApp() *App {
	cfg := config.Load()
	router := gin.New()
	lg := logger.New()

	// Middleware
	router.Use(logger.AddLoggerToGin(lg))
	router.Use(logger.GinLogger)
	router.Use(gin.Recovery())

	// Routes
	health.Register(router)

	return &App{
		Config: cfg,
		Router: router,
		Logger: lg,
	}
}

func (a *App) Start() {
	port := ":" + a.Config.Port
	router := a.Router
	lg := a.Logger

	if err := router.Run(port); err != nil {
		lg.Error("server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
