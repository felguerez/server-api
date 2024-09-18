package spotify

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"os"
	"time"
	"web-service/utils"
)

// ensureFreshTokens godoc
// @Summary Validates the token expiry and exchanges for new tokens if expired
func ensureFreshTokens(tokens *utils.Item) (string, string, time.Time, error) {
	if tokens == nil {
		return "", "", time.Now(), errors.New("nil tokens provided to ensureFreshToken")
	}
	accessToken := tokens.AccessToken
	refreshToken := tokens.RefreshToken
	expiresAt := time.Unix(tokens.ExpiresAt, 0)

	// access token is expired - use refresh token to get a new one
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
			return "", "", time.Time{}, err
		}

		accessToken = newToken.AccessToken
		refreshToken = newToken.RefreshToken
		expiresAt = newToken.Expiry
	}

	return accessToken, refreshToken, expiresAt, nil
}
