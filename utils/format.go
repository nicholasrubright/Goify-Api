package utils

import (
	b64 "encoding/base64"
	"fmt"
)



func GetEncodedClientCredentials(clientId string, clientSecret string) string {
	str := fmt.Sprintf("%s:%s", clientId, clientSecret)
	encodedStr := b64.StdEncoding.EncodeToString([]byte(str))
	return encodedStr
}

// Defaults to track
// Returns string in format: "spotify:track:[track id]"
func GetSpotifyURIsForAddingToPlaylist(track_id string) string {
	return fmt.Sprintf("spotify:track:%s", track_id)
}