package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TursoURL   string
	TursoToken string
}

var AppConfig Config

func LoadConfig() error {
	err := godotenv.Load("../on-page-seo/.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	AppConfig = Config{
		TursoURL:   os.Getenv("DATABASE_URL"),
		TursoToken: os.Getenv("AUTH_TOKEN"),
	}

	if AppConfig.TursoURL == "" || AppConfig.TursoToken == "" {
		return fmt.Errorf("Database credentials are not set in the environment variables")
	}

	return nil
}
