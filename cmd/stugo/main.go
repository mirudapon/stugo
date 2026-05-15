package main

import (
	"stugo/internal/app"
	"stugo/internal/config"
)

func main() {
	cfg := config.Load()

	app.Start(cfg)
}
