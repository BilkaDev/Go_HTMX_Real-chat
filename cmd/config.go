package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	State string
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

	config := Config{
		State: state,
		Port:  port,
	}
	return &config
}
