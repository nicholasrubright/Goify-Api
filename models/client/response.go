package models

import models "github.com/goify-api/models/spotify"

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

type ClientUserProfileResponse struct {
	Name	string				`json:"name"`
	Images	[]models.Images		`json:"images"`
}