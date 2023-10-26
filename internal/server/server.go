package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/phone-book/internal/config"
	"github.com/kozhamseitova/phone-book/internal/handler"
	"github.com/kozhamseitova/phone-book/pkg/logger"
)

type Server struct {
	app *fiber.App
}

func New(cfg *config.Config, logger logger.Logger, handler *handler.Handler) *Server {
	app := fiber.New()

	handler.InitRoutes(app)

	return &Server{
		app: app,
	}
}

func (s *Server) Run(port string) error {
	return s.app.Listen(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}