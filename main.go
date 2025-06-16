package main

import (
	"fmt"
	"net/http"

	"github.com/budsx/bookcabin/config"
	"github.com/budsx/bookcabin/controller"
	"github.com/budsx/bookcabin/repository"
	"github.com/budsx/bookcabin/services"
	"github.com/budsx/bookcabin/util/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	logger := logger.GetLogger()
	bookCabinService := services.NewBookCabinService(repo, logger)
	bookCabinController := controller.NewBookCabinController(bookCabinService)

	r.Get("/api/v1/seat-map", bookCabinController.GetSeatMap)
	// r.Post("/api/v1/seat-map/select", bookCabinController.SelectSeat)
	// r.Post("/api/v1/seat-map/confirm", bookCabinController.ConfirmSeat)

	// r.Get("/api/v1/seat-map/cancel", bookCabinController.CancelSeat)

	serverAddr := fmt.Sprintf(":%s", cfg.ServicePort)
	logger.Info(fmt.Sprintf("Server is running on port %s", cfg.ServicePort))
	http.ListenAndServe(serverAddr, r)
}
