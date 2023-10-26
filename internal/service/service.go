package service

import (
	"github.com/kozhamseitova/phone-book/internal/repository"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Service interface {

}

type service struct {
	repository repository.Repository
	logger logger.Logger
}

func New(repository repository.Repository, logger logger.Logger) Service {
	return &service{
		repository: repository,
		logger: logger,
	}
}