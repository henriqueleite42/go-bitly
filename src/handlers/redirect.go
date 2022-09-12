package handlers

import (
	"go-bitly/src/database/entities"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Redirect(ctx *fiber.Ctx) error {
	code := ctx.Params("code")

	shortURL := entities.ShortURLEntity{}
	err := entities.FindShortURLByCode(&shortURL, code)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting short url " + err.Error(),
		})
	}

	shortURL.Clicked += 1
	err = entities.UpdateShortURL(&shortURL)
	if err != nil {
		log.Printf("error updating %v", err.Error())
	}

	return ctx.Redirect(shortURL.Redirect, fiber.StatusTemporaryRedirect)
}
