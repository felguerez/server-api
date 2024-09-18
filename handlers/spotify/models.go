package spotify

import (
	"time"
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

type PlayableItem interface {
	GetName() string
	GetDuration() int
}

func (t Track) GetName() string {
	return t.Name
}

func (t Track) GetDuration() int {
	return t.Album.TotalTracks
}

func (e PodcastEpisode) GetName() string {
	return e.Name
}

func (e PodcastEpisode) GetDuration() int {
	return e.DurationMs
}

type RecentlyPlayedResponse struct {
	Cursors Cursors `json:"cursors"`
	Href    string  `json:"href"`
	Items   []Item  `json:"items"`
	Limit   int     `json:"limit"`
	Next    string  `json:"next"`
}

type CurrentlyPlayingResponse struct {
	Item                 interface{}  `json:"item,omitempty"` // Raw JSON for later manual unmarshalling
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

type PodcastEpisode struct {
	AudioPreviewURL      string       `json:"audio_preview_url"`
	Description          string       `json:"description"`
	DurationMs           int          `json:"duration_ms"`
	Explicit             bool         `json:"explicit"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	IsExternallyHosted   bool         `json:"is_externally_hosted"`
	IsPlayable           bool         `json:"is_playable"`
	Language             string       `json:"language"`
	Languages            []string     `json:"languages"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	ResumePoint          ResumePoint  `json:"resume_point"`
	Show                 PodcastShow  `json:"show"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
}

type PodcastShow struct {
	Name             string       `json:"name"`
	Publisher        string       `json:"publisher"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	Images           []Image      `json:"images"`
	URI              string       `json:"uri"`
	AvailableMarkets []string     `json:"available_markets"`
	Description      string       `json:"description"`
	HTMLDescription  string       `json:"html_description"`
}

type ResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMs int  `json:"resume_position_ms"`
}
