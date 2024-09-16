package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type PgConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Database string `env:"DB_NAME"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func NewConfig() PgConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return PgConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
