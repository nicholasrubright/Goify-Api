package main

// Token Response
type TokenResponse struct {
	Token	string	`json:"access_token"`
	Type	string	`json:"token_type"`
	Expires	int		`json:"expires_in"`
}

type ClientTokenResponse struct {
	Code	string	`json:"code"`
	Verifier	string	`json:"verifier"`
}

type ClientAuthResponse struct {
	Url	string	`json:"url"`
}