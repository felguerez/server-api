package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"web-service/utils"
)

type recentlyPlayedResponse struct {
	Items     *[]Track `json:"items"`
	IsPlaying bool     `json:"is_playing"`
	limit     int      `json:"limit"`
	href      string   `json:"href"`
	cursors   cursors  `json:"cursors"`
}

type cursors struct {
	after  string `json:"after"`
	before string `json:"before"`
}

// RecentlyPlayedTracks godoc
// @Summary Get recently played tracks by user from Spotify Web API
// @Description GET api.spotify.com/v1/me/player/recently-played. Sends back array of tracks.
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
	if tokens == nil {
		return c.Status(422).JSON(fiber.Map{"message": "Could not get item from dynamodb with key `felguerez` from RecentlyPlayedTracks"})
	}
	accessToken, _, _, err := ensureFreshTokens(tokens)
	if err != nil {
		return err
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
		fmt.Sprintf("we got an error: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	var recentlyPlayed RecentlyPlayedResponse
	err = json.NewDecoder(resp.Body).Decode(&recentlyPlayed)
	if err != nil {
		return err
	}
	return c.JSON(recentlyPlayed)
}
