package models

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