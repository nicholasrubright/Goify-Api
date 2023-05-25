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

// User's Playlists
type SpotifyCurrentUserPlaylistsResponse struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}

// User Profile
type SpotifyCurrentUserProfileResponse struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent ExplicitContent `json:"explicit_content"`
	ExternalUrls    ExternalUrls    `json:"external_urls"`
	Followers       Followers       `json:"followers"`
	Href            string          `json:"href"`
	ID              string          `json:"id"`
	Images          []Images        `json:"images"`
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
	Items    []PlaylistItems `json:"items"`
}


// Spotify Error Responses
type SpotifyAuthorizationErrorResponse struct {
	Error 			string	`json:"error"`
	Description 	string	`json:"error_description"`
}

type SpotifyErrorResponse struct {
	Error SpotifyError `json:"error"`
}


