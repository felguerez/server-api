package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"net/url"
	"web-service/utils"
)

// CurrentlyPlaying godoc
// @Summary Get Currently playing track from Spotify
// @Description Get the currently playing track or episode from Spotify.
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify/currently-playing [get]
func CurrentlyPlaying(c *fiber.Ctx) error {
	// Fetch tokens
	tokens, err := utils.GetItem("felguerez") // TODO: remove hardcoded key
	if err != nil {
		fmt.Println("Could not get item from dynamodb with key `felguerez`")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve tokens"})
	}

	// Ensure fresh access token
	accessToken, _, _, err := ensureFreshTokens(tokens)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to refresh access token"})
	}

	baseURL := "https://api.spotify.com/v1/me/player/currently-playing"
	params := url.Values{}
	params.Add("market", "us")
	params.Add("additional_types", "track,episode")

	// Append the query parameters to the base URL
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make API request to Spotify
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Failed to reach Spotify API"})
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusNoContent {
		return c.JSON(fiber.Map{
			"is_playing":             false,
			"item":                   nil,
			"currently_playing_type": "",
		})
	}

	var currentlyPlaying CurrentlyPlayingResponse
	if err := json.NewDecoder(resp.Body).Decode(&currentlyPlaying); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse Spotify response"})
	}

	// Handle dynamic unmarshalling for the "item" field based on the currently_playing_type
	switch currentlyPlaying.CurrentlyPlayingType {
	case "track":
		var track Track

		// Re-marshal Item back into JSON bytes
		itemBytes, err := json.Marshal(currentlyPlaying.Item)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to re-marshal track item"})
		}

		// Unmarshal into Track struct
		if err := json.Unmarshal(itemBytes, &track); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse track data"})
		}
		// Set the track into Item
		currentlyPlaying.Item = track

	case "episode":
		var episode PodcastEpisode

		// Re-marshal Item back into JSON bytes
		itemBytes, err := json.Marshal(currentlyPlaying.Item)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to re-marshal episode item"})
		}

		// Unmarshal into PodcastEpisode struct
		if err := json.Unmarshal(itemBytes, &episode); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse episode data"})
		}
		// Set the episode into Item
		currentlyPlaying.Item = episode

	default:
		// Handle case where currently_playing_type is unknown
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unknown currently playing type"})
	}

	// Return the full response
	return c.JSON(currentlyPlaying)
}
