package main

type UserPlaylists struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Items `json:"items"`
}
type ExternalUrls struct {
	Spotify string `json:"spotify"`
}
type Images struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type Owner struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}
type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type Items struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Images     `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        Tracks       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type UserProfile struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent ExplicitContent `json:"explicit_content"`
	ExternalUrls    ExternalUrls    `json:"external_urls"`
	Followers       Followers       `json:"followers"`
	Href            string          `json:"href"`
	ID              string          `json:"id"`
	Images          []Images        `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	URI             string          `json:"uri"`
}
type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}
