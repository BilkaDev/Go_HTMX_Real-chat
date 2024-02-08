package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	State     string
	JWTSecret string
}

func LoadEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Load env from file
	state := os.Getenv("STATE")
	if state == "" {
		log.Fatalf("ENV STATE must be defined")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("ENV PORT must be defined")
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatalf("ENV JWT_SECRET must be defined")
	}

	config := Config{
		State:     state,
		Port:      port,
		JWTSecret: JWTSecret,
	}
	return &config
}
