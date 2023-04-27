package router

import (
	"github.com/gofiber/fiber/v2"
	"web-service/handlers/spotify"
)

func AddSpotifyRoutes(router fiber.Router) {
	group := router.Group("/spotify")
	group.Get("/", spotify.ApiIndex)
	group.Get("/authenticate", spotify.BeginOAuth)
	group.Get("/callback", spotify.OAuthCallback)
	//group.Get("/recently-played", spotify.RecentlyPlayedTracks)
	//group.Get("/currently-playing", spotify.CurrentlyPlaying)
}
