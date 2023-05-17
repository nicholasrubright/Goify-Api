package main

// Client Requests (Meaning requests we get from the client)
type ClientTokenRequest struct {
	Code	string	`json:"code"`
}

type ClientProfileRequest struct {
	Token	string	`json:"token"`
}
