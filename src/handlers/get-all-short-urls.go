package handlers

import (
	"go-bitly/src/database/entities"

	"github.com/gofiber/fiber/v2"
)

func GetAllShortURLs(ctx *fiber.Ctx) error {
	shortURLs := []entities.ShortURLEntity{}
	err := entities.GetAllShortURLs(&shortURLs)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting short urls " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortURLs)
}
