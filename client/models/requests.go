package models
// Requests from the Client (Client = frontend)

// Token Request
type ClientTokenRequest struct {
	Code	string	`json:"code"`
}


// Create Playlist
type ClientCreatePlaylistRequest struct {
	User_ID		string		`json:"user_id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Playlists	[]string	`json:"playlists"`
}