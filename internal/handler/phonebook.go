package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/phone-book/api"
	"github.com/kozhamseitova/phone-book/internal/models"
)

func(h *Handler) search(c *fiber.Ctx) error {
	phone := c.Params("Search[phone]")
	name := c.Params("Search[name]")

	search := models.Search{
		Phone: phone,
		Name: name,
	}

	result, err := h.service.Search(c.UserContext(), search)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
		Data: result,
	})
}

func(h *Handler) create(c *fiber.Ctx) error {
	var req models.Search 

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = h.service.Create(c.UserContext(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
	})
}