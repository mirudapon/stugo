package main

import (
	"stugo/internal/app"
	"stugo/internal/config"
)

func main() {
	config.Load()

	cfg := config.Get()

	app.Start(cfg.Port, cfg.Mode)
}
