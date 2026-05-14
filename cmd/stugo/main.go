package main

import (
	"stugo/interval/app"
	"stugo/interval/config"
)

func main() {
	config.Load()

	cfg := config.Get()

	app.Start(cfg.Port, cfg.Mode)
}
