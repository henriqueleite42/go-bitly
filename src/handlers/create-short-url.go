package handlers

import (
	"go-bitly/src/database/entities"

	"github.com/gofiber/fiber/v2"
)

func CreateShortURL(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	shortURL := entities.ShortURLEntity{}
	err := ctx.BodyParser(&shortURL)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing json " + err.Error(),
		})
	}

	if err = entities.CreateShortURL(&shortURL); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not create short url " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(shortURL)
}
