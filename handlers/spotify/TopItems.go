package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"strings"
	"web-service/utils"
)

// TopItems godoc
// @Summary Get top items by type (artists or tracks)
// @Description GET https://api.spotify.com/v1/me/top/:type.
// @Description * Responds with `items` in JSON representing a list of artists or tracks.
// @Param type path string true "Type of item to get, either `artists` or `tracks`"
// @Param time_range query string false "Time range to query for top items, either `short_term`, `medium_term`, or `long_term` (default: `medium_term`)"
// @Tags spotify
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api/spotify/top/{type} [get]
func TopItems(c *fiber.Ctx) error {
	tokens, err := utils.GetItem("felguerez") // TODO: remove hardcoded key
	if err != nil {
		fmt.Println("Could not get item from dynamodb with key `felguerez`")
	}

	accessToken, _, _, err := ensureFreshTokens(tokens)
	if err != nil {
		return err
	}

	entity := c.Params("type")

	endpoint := fmt.Sprintf("https://api.spotify.com/v1/me/top/%s", entity)

	client := http.DefaultClient
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	q := req.URL.Query()
	q.Add("time_range", c.Query("time_range", "medium_term"))
	req.URL.RawQuery = q.Encode()
	fmt.Println(fmt.Sprintf("Getting %s from url: %s", entity, req.URL.String()))
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response, err := decodeResponse(c.Params("type"), resp.Body)
	if err != nil && strings.Contains(err.Error(), "key not found") {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err})
	} else if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Error decoding JSON response from API"})
	}
	return c.JSON(response)
}

func decodeResponse(thing string, body io.ReadCloser) (interface{}, error) {
	switch thing {
	case "artists":
		var topArtistsResponse TopArtistsResponse
		err := json.NewDecoder(body).Decode(&topArtistsResponse)
		if err != nil {
			return nil, err
		}
		return topArtistsResponse, nil
	case "tracks":
		var topTracksResponse TopTracksResponse
		err := json.NewDecoder(body).Decode(&topTracksResponse)
		if err != nil {
			return nil, err
		}
		return topTracksResponse, nil
	default:
		return nil, fmt.Errorf("key not found: %s", thing)
	}
}
