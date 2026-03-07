package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPListen string
	Env        string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		HTTPListen: getEnv("HTTP_LISTEN", ":8080"),
		Env:        getEnv("ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
