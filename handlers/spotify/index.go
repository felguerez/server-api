package spotify

import "time"

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
	Href         string       `json:"href"`
	Name         string       `json:"name"`
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