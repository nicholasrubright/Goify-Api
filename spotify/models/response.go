package models

// Responses from Spotify

// Token
type SpotifyTokenResponse struct {
	Token			string	`json:"access_token"`
	Type			string	`json:"token_type"`
	Scope			string	`json:"scope"`
	Expires			int		`json:"expires_in"`
	RefreshToken	string	`json:"refresh_token"`
}

// Authorization Url
type SpotifyAuthorizationUrl struct {
	Url		string		`json:"url"`
}

// User's Playlists
type SpotifyCurrentUserPlaylistsResponse struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []SpotifyItemsResponse `json:"items"`
}

// User Profile
type SpotifyCurrentUserProfileResponse struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent SpotifyExplicitContentResponse `json:"explicit_content"`
	ExternalUrls    SpotifyExternalUrlsResponse    `json:"external_urls"`
	Followers       SpotifyFollowersResponse       `json:"followers"`
	Href            string          `json:"href"`
	ID              string          `json:"id"`
	Images          []SpotifyImagesResponse        `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	URI             string          `json:"uri"`
}


// Playlists
type SpotifyPlaylistItemsResponse struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []SpotifyPlaylistItemsResponse `json:"items"`
}


// More responses

type SpotifyExternalUrlsResponse struct {
	Spotify string `json:"spotify"`
}
type SpotifyImagesResponse struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type SpotifyFollowersResponse struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type SpotifyOwnerResponse struct {
	ExternalUrls SpotifyExternalUrlsResponse `json:"external_urls"`
	Followers    SpotifyFollowersResponse    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}
type SpotifyTracksResponse struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type SpotifyItemsResponse struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  SpotifyExternalUrlsResponse `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []SpotifyImagesResponse     `json:"images"`
	Name          string       `json:"name"`
	Owner         SpotifyOwnerResponse        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        SpotifyTracksResponse       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type SpotifyExplicitContentResponse struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type SpotifyAddedByResponse struct {
	ExternalUrls SpotifyExternalUrlsResponse `json:"external_urls"`
	Followers    SpotifyFollowersResponse    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type SpotifyRestrictionsResponse struct {
	Reason string `json:"reason"`
}
type SpotifyCopyrightsResponse struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
type SpotifyExternalIdsResponse struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type SpotifyAlbumResponse struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         SpotifyExternalUrlsResponse `json:"external_urls"`
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
type SpotifyArtistsResponse struct {
	ExternalUrls SpotifyExternalUrlsResponse `json:"external_urls"`
	Followers    SpotifyFollowersResponse    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []SpotifyImagesResponse     `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
type SpotifyLinkedFromResponse struct {

}
type SpotifyTrackResponse struct {
	Album            SpotifyAlbumResponse        `json:"album"`
	Artists          []SpotifyArtistsResponse    `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      SpotifyExternalIdsResponse  `json:"external_ids"`
	ExternalUrls     SpotifyExternalUrlsResponse `json:"external_urls"`
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


// type SpotifyPlaylistItemsResponse struct {
// 	AddedAt string  `json:"added_at"`
// 	AddedBy SpotifyAddedByResponse `json:"added_by"`
// 	IsLocal bool    `json:"is_local"`
// 	Track   SpotifyTrackResponse   `json:"track"`
// }

type SpotifyCreatePlaylistResponse struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  SpotifyExternalUrlsResponse `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []SpotifyImagesResponse     `json:"images"`
	Name          string       `json:"name"`
	Owner         SpotifyOwnerResponse        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        SpotifyTracksResponse       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}


// Spotify Error Responses
type SpotifyAuthorizationErrorResponse struct {
	Error 			string	`json:"error"`
	Description 	string	`json:"error_description"`
}

type SpotifyErrorResponse struct {
	Error SpotifyError `json:"error"`
}

type SpotifyError struct {
	Status	int	`json:"status"`
	Message	string	`json:"message"`
}