package handlers

import (
	"github.com/gofiber/fiber/v2"
)

const SPOTIFY_CLIENT_ID = "SPOTIFY_CLIENT_ID"
const SPOTIFY_CLIENT_SECRET = "SPOTIFY_CLIENT_SECRET"
const SPOTIFY_REDIRECT_URI = "SPOTIFY_REDIRECT_URI"

// Index godoc
// @Summary Render an index.html page
// @Description Renders views/index.html with injected data
// @Tags root
// @Accept */*
// @Produce text/html
// @Success 200
// @Router / [get]
func Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Title": "yooooo",
		"hello": "greetings",
	})
}
