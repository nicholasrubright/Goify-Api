package models

import models "github.com/goify-api/models/spotify"

// Internal models that we will use for the server logic

// User Profile
type UserProfile struct {
	ID		string
	Name	string
	Images	[]models.Images	
}

// Playlist
type Playlist struct {
	ID			string
	Name		string
	Images		[]models.Images
	Description string
}