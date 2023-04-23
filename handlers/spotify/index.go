package spotify

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
	AlbumType            string       `json:"albumType"`
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

type Track struct {
	Name  string `json:"name"`
	Album Album  `json:"album"`
}
