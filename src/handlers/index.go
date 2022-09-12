package handlers

import "github.com/gofiber/fiber/v2"

func Setup(router *fiber.App) {
	router.Get("/short-urls", GetAllShortURLs)

	router.Get("/short-urls/:id", GetShortURL)

	router.Post("/short-urls", CreateShortURL)

	router.Patch("/short-urls/:id", UpdateShortURL)

	router.Delete("/short-urls/:id", DeleteShortURL)

	router.Get("/r/:code", Redirect)
}
