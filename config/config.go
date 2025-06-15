package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/budsx/bookcabin/util/logger"
)

type Config struct {
	ServicePort string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logger.WithError(err).Error("Failed to load .env file")
		return nil, err
	}

	return &Config{
		ServicePort: "8080",
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}, nil
}
