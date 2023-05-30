package models

// Requests to Spotify

// Create Playlists Request Body
type SpotifyCreatePlaylistRequest struct {
	Name	string	`json:"name"`
	Public	bool	`json:"public"`
	Description	string	`json:"description"`
}