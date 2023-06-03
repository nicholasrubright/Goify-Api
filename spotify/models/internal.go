package models

type PlaylistItems struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Item `json:"items"`
}

type Item struct {
	AddedAt string	`json:"added_at"`
	AddedBy	string	`json:"added_by"`
	IsLocal	bool	`json:"is_local"`
	Track	Track	`json:"track"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []SpotifyArtistsResponse    `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      SpotifyExternalIdsResponse  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       SpotifyLinkedFromResponse   `json:"linked_from"`
	Restrictions     SpotifyRestrictionsResponse `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []SpotifyImagesResponse     `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         SpotifyRestrictionsResponse `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Copyrights           []SpotifyCopyrightsResponse `json:"copyrights"`
	ExternalIds          SpotifyExternalIdsResponse  `json:"external_ids"`
	Genres               []string     `json:"genres"`
	Label                string       `json:"label"`
	Popularity           int          `json:"popularity"`
	AlbumGroup           string       `json:"album_group"`
	Artists              []SpotifyArtistsResponse    `json:"artists"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}