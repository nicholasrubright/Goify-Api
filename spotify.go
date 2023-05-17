package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)


const (
	API_URL = "https://api.spotify.com/v1"
	AUTHORIZE_URL = "https://accounts.spotify.com/authorize?"
	TOKEN_URL	= "https://accounts.spotify.com/api/token"
)


// Spotify API Endpoints
func getAccessToken(clientId string, code string, redirectUri string) (*ClientTokenResponse, *SpotifyAuthorizationErrorResponse, error) {

	formData := url.Values{
		"grant_type": {"authorization_code"},
		"code": {code},
		"redirect_uri": {redirectUri},
	}

	request, err := http.NewRequest("POST", TOKEN_URL, strings.NewReader(formData.Encode()))

	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authorizationKey := "Basic " + getEncodedClientCredentials()
	request.Header.Add("Authorization", authorizationKey)

	response, err := http.DefaultClient.Do(request)

	if checkStatus(response, err) != nil {
		fmt.Println("There was a problem making the request")

		var spotifyAuthorizationErrorResponse SpotifyAuthorizationErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyAuthorizationErrorResponse); err != nil {
			return nil, nil, err
		}

		fmt.Println("Error: ", spotifyAuthorizationErrorResponse)

		return nil, &spotifyAuthorizationErrorResponse, nil
	}

	var spotifyTokenResponse SpotifyTokenResponse 

	if err := json.NewDecoder(response.Body).Decode(&spotifyTokenResponse); err != nil {
		return nil, nil, err
	}

	return &ClientTokenResponse{
		Token: spotifyTokenResponse.Token,
	}, nil, nil
}



func getAuthUrl(clientId string, redirectUrl string) (*ClientAuthUrlResponse, error) {
	
	state := generateAuthState()

	params := url.Values{
		"client_id": {clientId},
		"response_type": {"code"},
		"redirect_uri": {redirectUrl},
		"scope": {"user-read-email user-read-private"},
		"state": {state},
	}

	authUrl := AUTHORIZE_URL + params.Encode()

	return &ClientAuthUrlResponse{
		Url: authUrl,
	}, nil
} 

func getUserProfile(token string) (*UserProfileResponse, *SpotifyErrorResponse, error) {

	endpointUrl := API_URL + "/v1/me"

	request, err := http.NewRequest("GET", endpointUrl, nil)

	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if checkStatus(response, err) != nil {
		fmt.Println("There was a problem making the request")

		var spotifyErrorResponse SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			return nil, nil, err
		}

		fmt.Println("Error: ", spotifyErrorResponse)
		return nil, &spotifyErrorResponse, nil
	}

	var userProfileResponse UserProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&userProfileResponse); err != nil {
		return nil, nil, err
	}

	return &userProfileResponse, nil, nil

}


// Get User's Playlists

// Create new Playlists

// Add Tracks to a Playlist
