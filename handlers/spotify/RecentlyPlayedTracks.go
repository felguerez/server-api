package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"
	"web-service/utils"
)

type recentlyPlayedResponse struct {
	Item      *Track   `json:"item"`
	Context   *Context `json:"context,omitempty"`
	IsPlaying bool     `json:"is_playing"`
}

// RecentlyPlayedTracks godoc
// @Summary Get recently played tracks by user from Spotify Web API
// @Description GET api.spotify.com/v1/me/player/recently-played
// @Description Sends back array of tracks
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify/recently-played [get]
func RecentlyPlayedTracks(c *fiber.Ctx) error {
	tokens, err := utils.GetItem("felguerez") // TODO: remove hardcoded key
	if err != nil {
		fmt.Println("Could not get item from dynamodb with key `felguerez`")
	}
	accessToken := tokens.AccessToken
	refreshToken := tokens.RefreshToken
	expiresAt := time.Unix(tokens.ExpiresAt, 0)

	// If access token has expired, use refresh token to get a new one
	if time.Now().After(expiresAt) {
		conf := &oauth2.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://accounts.spotify.com/api/token",
			},
		}

		token := &oauth2.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Expiry:       expiresAt,
		}

		newToken, err := conf.TokenSource(context.Background(), token).Token()
		if err != nil {
			return err
		}

		accessToken = newToken.AccessToken
		refreshToken = newToken.RefreshToken
		expiresAt = newToken.Expiry
	}

	// Make API request to get user's recently played tracks
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/recently-played", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var recentlyPlayed recentlyPlayedResponse
	err = json.NewDecoder(resp.Body).Decode(&recentlyPlayed)
	if err != nil {
		return err
	}

	return c.JSON(recentlyPlayed)
}
