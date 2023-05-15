package main

// Client Responses (Means responses we want to send to the client)
type ClientTokenResponse struct {
	Token	string	`json:"token"`
}

type ClientAuthUrlResponse struct {
	Url		string		`json:"url"`
	Verifier	string	`json:"verifier"`
}

// Spotify Responses (Responses from the spotify api)
type SpotifyTokenResponse struct {
	Token	string	`json:"access_token"`
	Type	string	`json:"token_type"`
	Expires	int		`json:"expires_in"`
}

type SpotifyAuthorizationErrorResponse struct {
	Error	string	`json:"error"`
	Description	string	`json:"error_description"`
}

type SpotifyErrorResponse struct {
	Error SpotifyError	`json:"error"`
}

type SpotifyError struct {
	Status	int		`json:"status"`
	Message	string	`json:"message"`
}