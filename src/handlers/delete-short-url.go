package handlers

import (
	"go-bitly/src/database/entities"

	"github.com/gofiber/fiber/v2"
)

func DeleteShortURL(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	err := entities.DeleteShortURL(ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting short url " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send([]byte{})
}
