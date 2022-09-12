package handlers

import (
	"go-bitly/src/database/entities"

	"github.com/gofiber/fiber/v2"
)

func GetShortURL(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	shortURL := entities.ShortURLEntity{}
	err := entities.GetShortURL(&shortURL, ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting short url " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortURL)
}
