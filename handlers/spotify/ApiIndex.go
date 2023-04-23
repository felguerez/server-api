package spotify

import "github.com/gofiber/fiber/v2"

// ApiIndex godoc
// @Summary Spotify API index route
// @Description Returns a version number
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify [get]
func ApiIndex(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"version": "1.0"})
}
