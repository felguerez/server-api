package router

import (
	"github.com/gofiber/fiber/v2"
	"web-service/handlers"
)

func AddSpotifyRoutes(router fiber.Router) {
	group := router.Group("/spotify")
	group.Get("/", handlers.SpotifyApiIndex)
	group.Get("/authenticate", handlers.SpotifyBeginOAuth)
	group.Get("/callback", handlers.SpotifyCallback)
}
