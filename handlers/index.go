package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"web-service/utils"
)

// Index godoc
// @Summary Render an index.html page
// @Description Renders Index view with variables injected
// @Tags root
// @Accept */*
// @Produce text/html
// @Success 200
// @Router / [get]
func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"Title": "yooooo",
		"hello": "greetings",
	})
}

// SpotifyApiIndex godoc
// @Summary API index route
// @Description Returns a version number
// @Tags api
// @Accept */*
// @Produce application/json
// @Success 200
// @Router /api [get]
func SpotifyApiIndex(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"version": "1.0"})
}

// SpotifyBeginOAuth godoc
// @Summary Begins Spotify OAuth token exchange for user to accept permissions
// @Description First step in the OAuth flow. Sets a cookie on `spotify_auth_state` (SpotifyStateKey) to read later, builds a URL with OAuth config in query params and redirects to the Spotify-hosted OAuth service
// @Tags spotify
// @Accept */*
// @Success 302
// @Router /api/spotify/authenticate [get]
func SpotifyBeginOAuth(ctx *fiber.Ctx) error {
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

const SPOTIFY_CLIENT_ID = "SPOTIFY_CLIENT_ID"
const SPOTIFY_CLIENT_SECRET = "SPOTIFY_CLIENT_SECRET"
const SPOTIFY_REDIRECT_URI = "SPOTIFY_REDIRECT_URI"

// SpotifyCallback godoc
// @Summary Uses the `req.query.code` sent after SpotifyBeginOauth for authorization_code flow
// @Description Following SpotifyBeginOAuth we get accessToken and refreshToken and write to db
// @Tags spotify
// @Accept */*
// @Success 200
// @Router /api/spotify/callback [get]
func SpotifyCallback(ctx *fiber.Ctx) error {
	//state := ctx.Cookies(utils.SpotifyStateKey)
	ctx.ClearCookie(utils.SpotifyStateKey)

	data := url.Values{
		"grant_type":   []string{"authorization_code"},
		"redirect_uri": []string{*utils.CopyString(os.Getenv(SPOTIFY_REDIRECT_URI))},
		"code":         []string{ctx.Query("code")},
	}.Encode()

	clientID := utils.CopyString(os.Getenv(SPOTIFY_CLIENT_ID))
	clientSecret := utils.CopyString(os.Getenv(SPOTIFY_CLIENT_SECRET))
	token := base64.StdEncoding.EncodeToString([]byte(*clientID + ":" + *clientSecret))
	authorization := "Basic " + token

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		tokenResponse := &utils.TokenResponse{}
		err := json.NewDecoder(resp.Body).Decode(tokenResponse)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		item := utils.Item{
			ExpiresAt:    time.Now().Add(3200 * time.Second).Unix(),
			AccessToken:  tokenResponse.AccessToken,
			RefreshToken: tokenResponse.RefreshToken,
			Id:           "felguerez",
		}
		fmt.Println(item)
		utils.PutItem(item)

		return ctx.JSON(tokenResponse)
	}

	return ctx.JSON(fiber.Map{"success": false})
}
