package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	api_scheme = "https"
	api_host = "api.spotify.com"
	api_version = "v1"

	AUTHORIZE_URL = "https://accounts.spotify.com/authorize"
	TOKEN_URL	= "https://accounts.spotify.com/api/token"
)


// Creates the Base API URL
func GetSpotifyAPIUrl(endpoint string) string {
	return fmt.Sprintf("%v://%v/%v/%v", api_scheme, api_host, api_version, endpoint)
}


// Checks the status of HTTP requetss
func CheckStatus(response *http.Response, err error) error {
	if err != nil {
		return err
	}

	statusOk := response.StatusCode >= 200 && response.StatusCode < 300

	if !statusOk {
		return errors.New("response was not successful")
	}

	return nil
}


func SendEndpointError(endpoint string, err error) {
	log.Println(fmt.Sprintln("%v Error: \n%v", endpoint, err))
}