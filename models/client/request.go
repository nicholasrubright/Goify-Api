package models

// Requests from the Client (Requests we get from the client)

// Token request
type ClientTokenRequest struct {
	Code	string	`json:"code"`
}