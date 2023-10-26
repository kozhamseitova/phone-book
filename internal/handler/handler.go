package handler

import (
	"github.com/kozhamseitova/phone-book/internal/service"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Handler struct {
	service service.Service
	logger logger.Logger
}

func New(service service.Service, logger logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger: logger,
	}
}