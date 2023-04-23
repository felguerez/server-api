package spotify

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
	"web-service/handlers"
	"web-service/utils"
)

// OAuthCallback godoc
// @Summary Uses the `req.query.code` sent after SpotifyBeginOauth for authorization_code flow
// @Description Following BeginOAuth we get accessToken and refreshToken and write to db
// @Tags spotify
// @Accept */*
// @Success 200
// @Router /api/spotify/callback [get]
func OAuthCallback(ctx *fiber.Ctx) error {
	//state := ctx.Cookies(utils.SpotifyStateKey)
	ctx.ClearCookie(utils.SpotifyStateKey)

	data := url.Values{
		"grant_type":   []string{"authorization_code"},
		"redirect_uri": []string{*utils.CopyString(os.Getenv(handlers.SPOTIFY_REDIRECT_URI))},
		"code":         []string{ctx.Query("code")},
	}.Encode()

	clientID := utils.CopyString(os.Getenv(handlers.SPOTIFY_CLIENT_ID))
	clientSecret := utils.CopyString(os.Getenv(handlers.SPOTIFY_CLIENT_SECRET))
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
