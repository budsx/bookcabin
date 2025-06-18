package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/budsx/bookcabin/config"
	"github.com/budsx/bookcabin/controller"
	"github.com/budsx/bookcabin/repository"
	"github.com/budsx/bookcabin/services"

	logger "github.com/budsx/bookcabin/util/logger"
	requestidmiddleware "github.com/budsx/bookcabin/util/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	logger := logger.GetLogger()
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

	r := chi.NewRouter()
	r.Use(requestidmiddleware.RequestIDMiddleware)
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

	bookCabinService := services.NewBookCabinService(repo, logger)
	bookCabinController := controller.NewBookCabinController(bookCabinService)

	r.Get("/api/v1/seat-map", bookCabinController.GetSeatMap)
	r.Post("/api/v1/seat-map/select", bookCabinController.SelectSeat)

	serverAddr := fmt.Sprintf(":%s", cfg.ServicePort)
	logger.Info(fmt.Sprintf("Server is running on port %s", cfg.ServicePort))

	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Error("Failed to start server")
		}
	}()

	GracefulShutdown(server, repo)
}

func GracefulShutdown(server *http.Server, repo *repository.Repository) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Shutdown signal received, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.WithError(err).Error("Server forced to shutdown")
	} else {
		logger.Info("HTTP server shutdown complete")
	}

	if err := repo.Close(); err != nil {
		logger.WithError(err).Error("Failed to close repository connections")
	} else {
		logger.Info("Database connections closed")
	}

	logger.Info("Server shutdown complete")
}
