package utils

import (
	"fmt"
	"math/rand"
	"net/url"
)

const (
	STATE_LENGTH = 16
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// Generate Code Verifier for Spotify API
func GenerateAuthorizationState() string {
	text := make([]rune, STATE_LENGTH)
	for i := range text {
		text[i] = LETTERS[rand.Intn(len(LETTERS))]
	}

	return string(text)
}

func GenerateAuthorizationUrl(parameters url.Values) string {
	return fmt.Sprintf("%v?%v", AUTHORIZE_URL, parameters.Encode())
}