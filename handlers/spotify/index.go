package spotify

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"os"
	"time"
	"web-service/utils"
)

type Context struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}

type Actions struct {
	Disallows Disallows `json:"disallows"`
}

type Disallows struct {
	InterruptingPlayback bool `json:"interrupting_playback"`
}

type Artist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Genres       []string     `json:"genres,omitempty"`
	Href         string       `json:"href"`
	Id           string       `json:"id,omitempty"`
	Images       []Image      `json:"images,omitempty"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity,omitempty"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}

type Image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"totalTracks"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	Id                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Type                 string       `json:"type"`
	Uri                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
}

type ExternalIds struct {
	Isrc string `json:"isrc"`
}

type Track struct {
	Name  string `json:"name"`
	Album Album  `json:"album"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

type DeviceObject struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

type TrackInfo struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      ExternalIds  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

type Item struct {
	Context   Context   `json:"context"`
	PlayedAt  time.Time `json:"played_at"`
	TrackInfo TrackInfo `json:"track"`
}

type RecentlyPlayedResponse struct {
	Cursors Cursors `json:"cursors"`
	Href    string  `json:"href"`
	Items   []Item  `json:"items"`
	Limit   int     `json:"limit"`
	Next    string  `json:"next"`
}

type CurrentlyPlayingResponse struct {
	Item                 *Track       `json:"item,omitempty"`
	Context              *Context     `json:"context,omitempty"`
	Timestamp            int64        `json:"timestamp"`
	ProgressMs           int          `json:"progress_ms"`
	IsPlaying            bool         `json:"is_playing"`
	ItemId               string       `json:"item_id"`
	Actions              Actions      `json:"actions"`
	CurrentlyPlayingType string       `json:"currently_playing_type"`
	Device               DeviceObject `json:"device"`
}

type TopArtistsResponse struct {
	Items    []Artist `json:"items"`
	Total    int      `json:"total"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
}

type TopTracksResponse struct {
	Items    []Track `json:"items"`
	Total    int     `json:"total"`
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
}

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

// TODO: determine if this is a good idea; I don't think so
func GroupConsecutiveTracks(recentlyPlayed *RecentlyPlayedResponse) [][]Item {
	var groupedItems [][]Item
	var tempGroup []Item

	for i := 0; i < len(recentlyPlayed.Items)-1; i++ {
		currentItem := recentlyPlayed.Items[i]
		nextItem := recentlyPlayed.Items[i+1]

		currentArtistID := currentItem.TrackInfo.Artists[0].Id
		nextArtistID := nextItem.TrackInfo.Artists[0].Id

		tempGroup = append(tempGroup, currentItem)

		if currentArtistID != nextArtistID {
			groupedItems = append(groupedItems, tempGroup)
			tempGroup = []Item{}
		}
	}
	tempGroup = append(tempGroup, recentlyPlayed.Items[len(recentlyPlayed.Items)-1])
	groupedItems = append(groupedItems, tempGroup)

	return groupedItems
}
