package main

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	CODE_VERIFER_LENGTH = 128
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// Generate Code Verifier
func generateCodeVerifier() string {
	text := make([]rune, CODE_VERIFER_LENGTH)
	for i := range text {
		text[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(text)
}

// Generate Code Challenge
func generateCodeChallenge(codeVerifier string) string {
	data := base64.StdEncoding.EncodeToString([]byte(codeVerifier))

	hash := sha256.New()
	hash.Write([]byte(data))

	digest := hash.Sum(nil)

	return string(digest)
} 
