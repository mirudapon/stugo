package config

import (
	"os"
	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	Port string
}

func Load() *Config {
	// .env.local overrides .env; godotenv.Load skips vars that already exist,
	// so loading .env.local first gives it priority.
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")

	cfg = &Config{
		Port: getEnv("APP_PORT", "8080"),
	}

	return cfg
}

func Get() *Config {
	if cfg == nil {
		Load()
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
