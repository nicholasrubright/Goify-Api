package models

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}
type Images struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type Owner struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}
type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type Items struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Images     `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        Tracks       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type SpotifyError struct {
	Status	int	`json:"status"`
	Message	string	`json:"message"`
}

type AddedBy struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}
type Copyrights struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
type ExternalIds struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Images     `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Copyrights           []Copyrights `json:"copyrights"`
	ExternalIds          ExternalIds  `json:"external_ids"`
	Genres               []string     `json:"genres"`
	Label                string       `json:"label"`
	Popularity           int          `json:"popularity"`
	AlbumGroup           string       `json:"album_group"`
	Artists              []Artists    `json:"artists"`
}
type Artists struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Images     `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
type LinkedFrom struct {
}
type Track struct {
	Album            Album        `json:"album"`
	Artists          []Artists    `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      ExternalIds  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}


type PlaylistItems struct {
	AddedAt string  `json:"added_at"`
	AddedBy AddedBy `json:"added_by"`
	IsLocal bool    `json:"is_local"`
	Track   Track   `json:"track"`
}