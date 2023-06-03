package models


// Token Response
type SpotifyTokenResponse struct {
	Token			string	`json:"access_token"`
	Type			string	`json:"token_type"`
	Scope			string	`json:"scope"`
	Expires			int		`json:"expires_in"`
	RefreshToken	string	`json:"refresh_token"`
}