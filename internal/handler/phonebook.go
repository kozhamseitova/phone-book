package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/phone-book/api"
	"github.com/kozhamseitova/phone-book/internal/models"
	"github.com/kozhamseitova/phone-book/utils"
)

func(h *Handler) search(c *fiber.Ctx) error {
	var search models.Search
	c.QueryParser(&search)

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

	if !checkParams(req) {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: utils.ErrInvalidParam.Error(),
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

func checkParams(req models.Search) bool {
	if req.Name == "" || len(req.Name) < 3 || req.Phone == "" || len(req.Phone) < 11 {
		return false
	}
	return true
	
}