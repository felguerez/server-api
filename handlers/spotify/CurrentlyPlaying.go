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
// @Description GET api.spotify.com/v1/me/player/currently-playing
// @Description Responds with
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api [get]
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

	if !currentlyPlaying.IsPlaying {
		return c.JSON(fiber.Map{"is_playing": false, "item": nil})
	}

	if err != nil {
		return err
	}
	return c.JSON(currentlyPlaying)
}
