package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/goify-api/spotify/models"
	"github.com/goify-api/utils"
)

const (
	GRANT_TYPE = "authorization_code"
	SCOPE = "user-read-private user-read-email playlist-modify-public playlist-modify-private"
)

func GetAuthorizationUrl(clientId string, redirectUri string) (*models.SpotifyAuthorizationUrl, error) {
	state := utils.GenerateAuthorizationState()

	parameters := url.Values{
		"client_id": {clientId},
		"response_type": {"code"},
		"redirect_uri": {redirectUri},
		"scope": {SCOPE},
		"state": {state},
		"show_dialog": {"true"},
	}

	authorizationUrl := utils.GenerateAuthorizationUrl(parameters)

	return &models.SpotifyAuthorizationUrl{
		Url: authorizationUrl,
	}, nil
}

func GetAccessToken(clientId string, clientSecret string, code string, redirectUri string) (*models.SpotifyTokenResponse, *models.SpotifyAuthorizationErrorResponse, error) {

	formData := url.Values{
		"grant_type": {GRANT_TYPE},
		"code": {code},
		"redirect_uri": {redirectUri},
	}


	// Create request
	request, err := http.NewRequest("POST", utils.TOKEN_URL, strings.NewReader(formData.Encode()))


	if err != nil {
		utils.SendEndpointError("GetAccessToken:Request", err)
		return nil, nil, err 
	}

	// Headers
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authorizationKey := "Basic " + utils.GetEncodedClientCredentials(clientId, clientSecret)
	request.Header.Add("Authorization", authorizationKey)


	// Parsing Response
	response, err := http.DefaultClient.Do(request)

	if utils.CheckStatus(response, err) != nil {
		var spotifyAuthorizationErrorResponse models.SpotifyAuthorizationErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyAuthorizationErrorResponse); err != nil {
			utils.SendEndpointError("GetAccessToken:Decode", err)
			return nil, nil, err
		}

		utils.SendEndpointError("GetAccessToken:Status", errors.New(spotifyAuthorizationErrorResponse.Description))

		return nil , &spotifyAuthorizationErrorResponse, nil
		
	}

	var spotifyTokenResponse models.SpotifyTokenResponse

	if err := json.NewDecoder(response.Body).Decode(&spotifyTokenResponse); err != nil {
		utils.SendEndpointError("GetAccessToken:Decode", err)
		return nil, nil, err
	}

	return &spotifyTokenResponse, nil, nil
}

func GetCurrentUserProfile(token string) (*models.SpotifyCurrentUserProfileResponse, *models.SpotifyErrorResponse, error) {

	url := utils.GetSpotifyAPIUrl("me")

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		utils.SendEndpointError("GetCurrentUserProfile:Request", err)
		return nil, nil, err
	}

	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if utils.CheckStatus(response, err) != nil {
		var spotifyErrorResponse models.SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			utils.SendEndpointError("GetCurrentUserProfile:Decode", err)
			return nil, nil, err
		}

		utils.SendEndpointError("GetCurrentUserProfile:Status", errors.New(spotifyErrorResponse.Error.Message))

		return nil , &spotifyErrorResponse, nil
	}

	var spotifyCurrentUserProfileResponse models.SpotifyCurrentUserProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&spotifyCurrentUserProfileResponse); err != nil {
		utils.SendEndpointError("GetCurrentUserProfile:Decode", err)
		return nil, nil, err
	}

	return &spotifyCurrentUserProfileResponse, nil, nil
}

func GetCurrentUserPlaylists(token string) (*models.SpotifyCurrentUserPlaylistsResponse, *models.SpotifyErrorResponse, error) {
	url := utils.GetSpotifyAPIUrl("me/playlists")

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		utils.SendEndpointError("GetCurrentUserPlaylists:Request", err)
		return nil, nil, err
	}

	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if utils.CheckStatus(response, err) != nil {
		var spotifyErrorResponse models.SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			utils.SendEndpointError("GetCurrentUserPlaylists:Decode", err)
			return nil, nil, err
		}

		utils.SendEndpointError("GetCurrentUserPlaylists:Status", errors.New(spotifyErrorResponse.Error.Message))

		return nil , &spotifyErrorResponse, nil
	}

	var spotifyCurrentUserPlaylistsResponse models.SpotifyCurrentUserPlaylistsResponse

	if err := json.NewDecoder(response.Body).Decode(&spotifyCurrentUserPlaylistsResponse); err != nil {
		utils.SendEndpointError("GetCurrentUserPlaylists:Decode", err)
		return nil, nil, err
	}

	return &spotifyCurrentUserPlaylistsResponse, nil, nil
}

func CreatePlaylistForUser(userId string, playlist_name string, playlist_description string, token string) (*models.SpotifyCreatePlaylistResponse, *models.SpotifyErrorResponse) {
	url := utils.GetSpotifyAPIUrl(fmt.Sprintln("users/%v/playlists", userId))

	spotifyCreatePlaylistRequest := models.SpotifyCreatePlaylistRequest {
		Name: playlist_name,
		Description: playlist_description,
		Public: false,
	}

	requestBody, err := json.Marshal(spotifyCreatePlaylistRequest)

	if err != nil {
		utils.SendEndpointError("CreatePlaylists:Decode", err)
		return nil, &models.SpotifyErrorResponse{
			Error: models.SpotifyError{
				Status: 500,
				Message: err.Error(),
			},
		}
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil {
		utils.SendEndpointError("CreatePlaylists:Request", err)
		return nil, &models.SpotifyErrorResponse{
			Error: models.SpotifyError{
				Status: 500,
				Message: err.Error(),
			},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if utils.CheckStatus(response, err) != nil {

		utils.SendEndpointError("CreatePlaylists:Status", err)

		var spotifyErrorResponse models.SpotifyErrorResponse

		if err := json.NewDecoder(response.Body).Decode(&spotifyErrorResponse); err != nil {
			return nil, &models.SpotifyErrorResponse{
				Error: models.SpotifyError{
					Status: 500,
					Message: err.Error(),
				},
			}
		}

		return nil, &spotifyErrorResponse
	}

	var spotifyCreatePlaylistResponse models.SpotifyCreatePlaylistResponse

	if err := json.NewDecoder(response.Body).Decode(&spotifyCreatePlaylistResponse); err != nil {
		return nil, &models.SpotifyErrorResponse{
			Error: models.SpotifyError{
				Status: 500,
				Message: err.Error(),
			},
		}
	}

	

	return &spotifyCreatePlaylistResponse, nil
}

func GetTracksFromPlaylist(playlist_id string, token string) (*models.SpotifyPlaylistItemsResponse, *models.SpotifyErrorResponse) {
	return nil, nil
}

func AddTrackToPlaylist(playlist_id string, track_ids []string, token string) {

}