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
// @Summary Redirects to spotify.com to begin OAuth2 token exchange
// @Description Sets a cookie with the state key, builds
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

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func SpotifyCallback(ctx *fiber.Ctx) error {
	//state := ctx.Cookies(utils.SpotifyStateKey)
	ctx.ClearCookie(utils.SpotifyStateKey)

	data := url.Values{
		"grant_type":   []string{"client_credentials"},
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
		tokenResponse := &TokenResponse{}
		err := json.NewDecoder(resp.Body).Decode(tokenResponse)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		//item := map[string]interface{}{
		//	"expiresAt":    time.Now().Add(3200 * time.Second).Unix(),
		//	"accessToken":  tokenResponse.AccessToken,
		//	"refreshToken": tokenResponse.RefreshToken,
		//	"id":           "felguerez",
		//}

		//tableName := utils.CopyString(os.Getenv("TABLE_NAME"))
		//_, err = Dynamo.Put(tableName, item).Run()
		//if err != nil {
		//	return ctx.SendStatus(fiber.StatusInternalServerError)
		//}

		return ctx.JSON(tokenResponse)
	}

	return ctx.JSON(fiber.Map{"success": false})
}
