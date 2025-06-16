package main

import (
	"net/http"

	"github.com/budsx/bookcabin/config"
	"github.com/budsx/bookcabin/controller"
	"github.com/budsx/bookcabin/repository"
	"github.com/budsx/bookcabin/services"
	"github.com/budsx/bookcabin/util/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	logger := logger.GetLogger()
	bookCabinService := services.NewBookCabinService(repo, logger)
	bookCabinController := controller.NewBookCabinController(bookCabinService)

	r.Get("/api/v1/seat-map", bookCabinController.GetSeatMap)
	// TODO: select seat


	logger.Info("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
