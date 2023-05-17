package main

import (
	b64 "encoding/base64"
	"errors"
	"math/rand"
	"net/http"
)

const (
	STATE_LENGTH = 16
)

var (
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// Generate Code Verifier
func generateAuthState() string {
	text := make([]rune, STATE_LENGTH)
	for i := range text {
		text[i] = LETTERS[rand.Intn(len(LETTERS))]
	}

	return string(text)
}

func getEncodedClientCredentials() string {
	str := CLIENT_ID + ":" + CLIENT_SECRET
	encodedStr := b64.StdEncoding.EncodeToString([]byte(str))

	return encodedStr
}



// check Status of HTTP requests
func checkStatus(response *http.Response, err error) error {

	if err != nil {
		return err
	}

	statusOk := response.StatusCode >= 200 && response.StatusCode < 300

	if !statusOk {
		return errors.New("response was not successful")
	}

	return nil

}