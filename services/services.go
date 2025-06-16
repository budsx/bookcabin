package services

import (
	"github.com/budsx/bookcabin/repository"
	"go.uber.org/zap"
)

type BookCabinService struct {
	repo   *repository.Repository
	logger *zap.Logger
}

func NewBookCabinService(repo *repository.Repository, logger *zap.Logger) *BookCabinService {
	return &BookCabinService{repo: repo, logger: logger}
}
