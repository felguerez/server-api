package spotify

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"os"
	"web-service/utils"
)

// BeginOAuth godoc
// @Summary Begins Spotify OAuth token exchange for user to accept permissions
// @Description First step in the OAuth flow. Sets a cookie on `spotify_auth_state` (SpotifyStateKey) to read later, builds a URL with OAuth config in query params and redirects to the Spotify-hosted OAuth service
// @Tags spotify
// @Accept */*
// @Success 302
// @Router /api/spotify/authenticate [get]
func BeginOAuth(ctx *fiber.Ctx) error {
	state := utils.RandStringBytes(16)
	ctx.Cookie(&fiber.Cookie{Name: utils.SpotifyStateKey, Value: state})
	query := url.Values{}
	query.Set("response_type", "code")
	query.Set("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	query.Set("scope", utils.Scope)
	query.Set("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))
	query.Set("state", state)

	spotifyAuthURL := fmt.Sprintf("%s?%s", "https://accounts.spotify.com/authorize", query.Encode())
	return ctx.Redirect(spotifyAuthURL)
}
