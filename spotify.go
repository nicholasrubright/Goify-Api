package main

import (
	"bytes"
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

	SCOPE = "user-read-private user-read-email playlist-modify-public playlist-modify-private"
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


		// if spotifyAuthorizationErrorResponse.Error == "invalid_grant" && past_token != "" {
		// 	token, err := getRefreshToken(past_token)
		// }


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
		"show_dialog": {"true"},
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
		ID: userProfileResponse.ID,
	}

	return clientProfileResponse, nil, nil

}



func getUserPlaylists(token string) (*clientModels.ClientUserPlaylistsResponse, *spotifyModels.SpotifyErrorResponse, error) {
	
	endpointUrl := API_URL + "/me/playlists"

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

	var userPlaylistsResponse spotifyModels.SpotifyCurrentUserPlaylistsResponse

	if err:= json.NewDecoder(response.Body).Decode(&userPlaylistsResponse); err != nil {
		return nil, nil, err
	}

	clientUserPlaylistResponse, _ := transformPlaylist(userPlaylistsResponse)




	return clientUserPlaylistResponse, nil, nil
}



// Turn playlists into response
func transformPlaylist(userplaylistResponse spotifyModels.SpotifyCurrentUserPlaylistsResponse) (*clientModels.ClientUserPlaylistsResponse, error) {

	playlistResponse := clientModels.ClientUserPlaylistsResponse{
		Limit: userplaylistResponse.Limit,
		Next: userplaylistResponse.Next,
		Offset: userplaylistResponse.Offset,
		Previous: userplaylistResponse.Previous,
		Total: userplaylistResponse.Total,
	}

	var playlists []clientModels.ClientPlaylistResponse

	for _, item := range userplaylistResponse.Items {
		playlists = append(playlists, clientModels.ClientPlaylistResponse{
			ID: item.ID,
			Images: item.Images,
			Name: item.Name,
		})
	}

	playlistResponse.Playlists = playlists

	return &playlistResponse, nil

}

// Create new Playlists
func createPlaylist(userId string, name string, description string, token string) (*spotifyModels.SpotifyErrorResponse, error) {

	endpointUrl := API_URL + "/users/" + userId + "/playlists"

	playlistRequest := spotifyModels.SpotifyCreatePlaylistRequest{
		Name: name,
		Description: description,
		Public: false,
	}

	requestBody, err := json.Marshal(playlistRequest)

	if err != nil {
		fmt.Println("Could not marshal playlist request")
		return nil, err
	}


	request, err := http.NewRequest("POST", endpointUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		fmt.Println("Could not make post requet for playlist request")
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if checkStatus(response, err) != nil {
		fmt.Println("There was a problem making the request")

		var spotifyErrorResponse spotifyModels.SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			return nil, err
		}

		fmt.Println("Error: ", spotifyErrorResponse)
		return &spotifyErrorResponse, nil
	}

	var createPlaylistResponse spotifyModels.Items

	if err := json.NewDecoder(response.Body).Decode(&createPlaylistResponse); err != nil {
		return nil, err
	}

	fmt.Println("Response: ", createPlaylistResponse)

	return nil, nil
}



// Add Tracks to a Playlist
