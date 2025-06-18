package services

import (
	"github.com/budsx/bookcabin/repository"
	"github.com/sirupsen/logrus"
)

type BookCabinService struct {
	repo   *repository.Repository
	logger *logrus.Logger
}

func NewBookCabinService(repo *repository.Repository, logger *logrus.Logger) *BookCabinService {
	return &BookCabinService{repo: repo, logger: logger}
}
