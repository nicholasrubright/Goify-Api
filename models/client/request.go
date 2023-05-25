package models

// Requests from the Client (Requests we get from the client)

// Token request
type ClientTokenRequest struct {
	Code	string	`json:"code"`
}


type ClientCreatePlaylistResponse struct {
	User		string		`json:"user"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Playlists	[]string	`json:"playlists"`
}