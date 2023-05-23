package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	clientModels "github.com/goify-api/models/client"
	spotifyModels "github.com/goify-api/models/spotify"
)


const (
	API_URL = "https://api.spotify.com/v1"
	AUTHORIZE_URL = "https://accounts.spotify.com/authorize?"
	TOKEN_URL	= "https://accounts.spotify.com/api/token"

	SCOPE = "user-read-private user-read-email"
)


// Spotify API Endpoints
func getAccessToken(clientId string, code string, redirectUri string) (*clientModels.ClientTokenResponse, *spotifyModels.SpotifyAuthorizationErrorResponse, error) {

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

		var spotifyAuthorizationErrorResponse spotifyModels.SpotifyAuthorizationErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyAuthorizationErrorResponse); err != nil {
			return nil, nil, err
		}

		fmt.Println("Error: ", spotifyAuthorizationErrorResponse)

		return nil, &spotifyAuthorizationErrorResponse, nil
	}

	var spotifyTokenResponse spotifyModels.SpotifyTokenResponse 

	if err := json.NewDecoder(response.Body).Decode(&spotifyTokenResponse); err != nil {
		return nil, nil, err
	}

	return &clientModels.ClientTokenResponse{
		Token: spotifyTokenResponse.Token,
	}, nil, nil
}



func getAuthUrl(clientId string, redirectUrl string) (*clientModels.ClientAuthorizationUrlResponse, error) {
	
	state := generateAuthState()

	params := url.Values{
		"client_id": {clientId},
		"response_type": {"code"},
		"redirect_uri": {redirectUrl},
		"scope": {SCOPE},
		"state": {state},
	}

	authUrl := AUTHORIZE_URL + params.Encode()

	return &clientModels.ClientAuthorizationUrlResponse{
		Url: authUrl,
	}, nil
} 

func getUserProfile(token string) (*clientModels.ClientUserProfileResponse, *spotifyModels.SpotifyErrorResponse, error) {

	endpointUrl := API_URL + "/me"

	request, err := http.NewRequest("GET", endpointUrl, nil)

	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if checkStatus(response, err) != nil {
		fmt.Println("There was a problem making the request")

		var spotifyErrorResponse spotifyModels.SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			return nil, nil, err
		}

		fmt.Println("Error: ", spotifyErrorResponse)
		return nil, &spotifyErrorResponse, nil
	}

	var userProfileResponse spotifyModels.SpotifyCurrentUserProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&userProfileResponse); err != nil {
		return nil, nil, err
	}


	clientProfileResponse := &clientModels.ClientUserProfileResponse{
		Name: userProfileResponse.DisplayName,
		Images: userProfileResponse.Images,
	}

	return clientProfileResponse, nil, nil

}


func getUserPlaylists(token string) (*spotifyModels.SpotifyCurrentUserPlaylistsResponse, *spotifyModels.SpotifyErrorResponse, error) {
	// endpointUrl := API_URL

	// request, err := http.NewRequest("GET", endpointUrl, nil)

	// if err != nil {
	// 	return nil, nil, err
	// }

	// request.Header.Add("Authorization", "Bearer " + token)

	return nil, nil, nil
}

// Get User's Playlists

// Create new Playlists

// Add Tracks to a Playlist
