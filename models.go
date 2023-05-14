package main

// User
type UserProfile struct {
	Country		string		`json:"country"`
	DisplayName	string		`json:"display_name"`
	Email		string		`json:"email"`
	ExplicitContent	ExplicitContent	`json:"explicit_content"`
	ExternalUrls	ExternalUrls	`json:"external_urls"`
	Followers		Followers		`json:"followers"`
	Href			string			`json:"href"`
	Id				string			`json:"id"`
	Images			[]Image			`json:"images"`
	Product			string			`json:"product"`
	Type			string			`json:"type"`
	Uri				string			`json:"uri"`
}

type Image struct {
	Url		string		`json:"url"`
	Height	int		`json:"height"`
	Width	int		`json:"width"`
}

type Followers struct {
	Href	string	`json:"href"`
	Total	int		`json:"total"`
}

type ExternalUrls struct {
	Spotify	string	`json:"spotify"`
}

type ExplicitContent struct {
	FilterEnabled	bool	`json:"filter_enabled"`
	FilterLocked	bool	`json:"filter_locked"`
}