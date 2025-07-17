package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
	SERVER_PORT  string
}

func LoadConfig(path string) *Config {
	_ = godotenv.Load(path)

	cfg := &Config{
		SERVER_PORT:  getEnvOrPanic("SERVER_PORT"),
		DATABASE_URL: getEnvOrPanic("DATABASE_URL"),
	}

	return cfg
}

func getEnvOrPanic(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		panic(fmt.Sprintf("❌ Required environment variable %q is missing or empty", key))
	}
	return val
}
