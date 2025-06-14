package main

import (
	"bookcabin/config"
	"bookcabin/repository"
	"bookcabin/util/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load config")
		return
	}

	repo, err := repository.NewBookCabinRepository(&repository.RepoConfig{
		DBConfig: repository.DBConfig{
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			DBName:   cfg.DBName,
		},
	})
	if err != nil {
		logger.WithError(err).Fatal("Failed to create repository")
		return
	}
	defer repo.Close()

}
