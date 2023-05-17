package main

// User
type UserProfileResponse struct {
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


type PlaylistResponse struct {
	Collaborative	bool	`json:"collaborative"`
	Description		string	`json:"description"`
	ExternalUrls	ExternalUrls	`json:"external_urls"`
	Followers		Followers		`json:"followers"`
	Href			string			`json:"href"`
	Id				string			`json:"id"`
	Images			[]Image			`json:"images"`
	Name			string			`json:"name"`
	Owner			Owner			`json:"owner"`
	Public			bool			`json:"public"`
	SnapshotId		string			`json:"snapshot_id"`

}

type Owner struct {
	ExternalUrls	ExternalUrls	`json:"external_urls"`
	Followers		Followers		`json:"followers"`
	Href			string			`json:"href"`
	Id				string			`json:"id"`
	Type			string			`json:"type"`
	Uri				string			`json:"uri"`
	DisplayName		string			`json:"display_name"`
}

type Tracks struct {
	Href	string	`json:"href"`
	Limit	int		`json:"limit"`
	Next	string	`json:"next"`
	Offset	int		`json:"offset"`
	Previous	string	`json:"previous"`
	Total	int		`json:"total"`
	Items	[]Item	`json:"items"`
}

type Item struct {
	AddedAt	string	`json:"added_at"`
	AddedBy	AddedBy	`json:"added_by"`
	IsLocal	bool	`json:"is_local"`
	Track	
}

type Track struct {
	Album	
}

type Album struct {
	AlbumType	string	`json:"album_type"`
	TotalTracks	int		`json:"total_tracks"`
}

type AddedBy struct {
	ExternalUrls	ExternalUrls	`json:"external_urls"`
	Followers		Followers		`json:"followers"`
	Href			string			`json:"href"`
	Id				string			`json:"id"`
	Type			string			`json:"type"`
	Uri				string			`json:"uri"`
}