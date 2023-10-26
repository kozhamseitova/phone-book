package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) InitRoutes(router *fiber.App) {

	api := router.Group("/api")
	api.Use(h.generateTraceId)

	v1 := api.Group("/v1")

	lead := v1.Group("/lead")
	lead.Get("/")
	lead.Post("/")
}