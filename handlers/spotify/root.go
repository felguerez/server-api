package spotify

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// ApiRoot godoc
// @Summary Spotify API index route
// @Description Returns a version number
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify [get]
func ApiRoot(ctx *fiber.Ctx) error {
	fmt.Println("we're hitting the index")
	return ctx.JSON(fiber.Map{"version": "1.0"})
}
