package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
 

type Config struct{
	DBUrl  string
    Port   string
    JwtKey string
}

func LoadConfig() *Config{
	_=godotenv.Load()
	cfg := &Config{
        DBUrl:  os.Getenv("DATABASE_URL"),
        Port:   os.Getenv("PORT"),
        JwtKey: os.Getenv("JWT_SECRET"),
    }
	if cfg.DBUrl == "" || cfg.JwtKey == "" {
        log.Fatal("missing critical environment variables")
    }

    if cfg.Port == "" {
        cfg.Port = "8080"
    }

	return cfg
}