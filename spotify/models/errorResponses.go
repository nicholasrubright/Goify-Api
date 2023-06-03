package models

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