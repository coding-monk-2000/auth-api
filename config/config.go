package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port      string
	DBDriver  string
	JWTSecret string
}

func NewFromEnv() (Config, error) {
	c := Config{
		Port:      os.Getenv("PORT"),
		DBDriver:  os.Getenv("DB_DRIVER"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
	if c.Port == "" || c.DBDriver == "" || c.JWTSecret == "" {
		return Config{}, fmt.Errorf("missing required env var")
	}
	return c, nil
}
