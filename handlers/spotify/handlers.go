package spotify

import (
	"github.com/gofiber/fiber/v2"
)

func AddSpotifyRoutes(router fiber.Router) {
	group := router.Group("/spotify")
	group.Get("/", ApiRoot)
	group.Get("/authenticate", BeginOAuth)
	group.Get("/callback", OAuthCallback)
	group.Get("/recently-played", RecentlyPlayedTracks)
	group.Get("/currently-playing", CurrentlyPlaying)
	group.Get("/top/:type", TopItems)
}
