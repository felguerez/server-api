package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"web-service/utils"
)

// CurrentlyPlaying godoc
// @Summary Get Currently playing track from Spotify
// @Description GET https://api.spotify.com/v1/me/player/currently-playing
// @Description * For currently playing music tracks, this endpoint responds in JSON with a currently playing `Track` as `item`.
// @Description * For currently playing podcasts, this endpoint responds in JSON with `{ "is_playing": "true, "item", nil, "currently_playing_type": "episode" }`. Spotify's API doesn't provide any episode data.
// @Description * When not currently listening, this endpoint responds in JSON with `{ "is_playing": false, "item": nil }`.
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify/currently-playing [get]
func CurrentlyPlaying(c *fiber.Ctx) error {
	tokens, err := utils.GetItem("felguerez") // TODO: remove hardcoded key
	if err != nil {
		fmt.Println("Could not get item from dynamodb with key `felguerez`")
	}
	accessToken, _, _, err := ensureFreshTokens(tokens)
	if err != nil {
		return err
	}
	// Make API request to get user's recently played tracks
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/currently-playing", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var currentlyPlaying CurrentlyPlayingResponse
	err = json.NewDecoder(resp.Body).Decode(&currentlyPlaying)

	if currentlyPlaying.CurrentlyPlayingType == "episode" {
		return c.JSON(fiber.Map{"is_playing": true, "item": nil, "currently_playing_type": currentlyPlaying.CurrentlyPlayingType})
	}

	if !currentlyPlaying.IsPlaying {
		return c.JSON(fiber.Map{"is_playing": false, "item": nil})
	}

	if err != nil {
		return err
	}
	return c.JSON(currentlyPlaying)
}
