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

	servicePort := os.Getenv("SERVICE_PORT")
	if servicePort == "" {
		servicePort = "8080"
	}

	return &Config{
		ServicePort: servicePort,
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}, nil
}
