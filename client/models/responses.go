package models

// Responses to the Client

// Authorization URL Model
type ClientAuthorizationUrlResponse struct {
	Url	string	`json:"url"`
}

// Token Response
type ClientTokenResponse struct {
	Token		string	`json:"token"`
	ExpiresIn	int		`json:"expires_in"`
}


// User Profile Response
type ClientUserProfileResponse struct {
	ID		string				`json:"id"`
	Name	string				`json:"name"`
	Images	[]ClientImageResponse		`json:"images"`
}


// User Playlist Response
type ClientUserPlaylistsResponse struct {
	Limit	int		`json:"limit"`
	Next	string	`json:"next"`
	Offset	int		`json:"offset"`
	Previous	string	`json:"previous"`
	Total	int		`json:"total"`
	Playlists	[]ClientPlaylistResponse	`json:"playlists"`
}

type ClientPlaylistResponse struct {
	ID	string	`json:"id"`
	Images	[]ClientImageResponse	`json:"images"`
	Name	string	`json:"name"`
}


type ClientImageResponse struct {
	URL	string	`json:"url"`
	Height	int	`json:"height"`
	Width	int `json:"width"`
}


// Create Playlist Response
type ClientCreatePlaylistResponse struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
	Tracks	[]ClientCreatePlaylistTrackResponse	`json:"tracks"`
}

type ClientCreatePlaylistTrackResponse struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
}

type ClientCombinePlaylistResponse struct {
	ID string	`json:"id"`
	Name	string	`json:"name"`
}

// Error Responses for Client

type ClientErrorResponse struct {
	Error	string	`json:"error"`
	Message string	`json:"message"`
}