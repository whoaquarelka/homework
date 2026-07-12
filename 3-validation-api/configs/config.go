package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Config EmailCred
}

type EmailCred struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *EmailConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file, using default config")
	}

	return &EmailConfig{
		EmailCred{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Address:  os.Getenv("EMAIL_ADDRESS")},
	}
}
