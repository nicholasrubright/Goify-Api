package main

import (
	"encoding/json"
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

func getAccessToken(clientId string, code string, verifier string, redirectUri string) (string, error) {

	formData := url.Values{
		"client_id": {clientId},
		"grant_type": {"authorization_code"},
		"code": {code},
		"redirect_uri": {redirectUri},
		"code_verifier": {verifier},
	}

	request, err := http.NewRequest("POST", TOKEN_URL, strings.NewReader(formData.Encode()))

	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	var tokenResponse TokenResponse 

	if err := json.NewDecoder(response.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}



func getAuthUrl(clientId string, redirectUrl string) (string, error) {
	
	verifier := generateCodeVerifier()
	challenge := generateCodeChallenge(verifier)

	params := url.Values{
		"client_id": {clientId},
		"response_type": {"code"},
		"redirect_uri": {redirectUrl},
		"scope": {"user-read-email user-read-private"},
		"code_challenge_method": {"S256"},
		"code_challenge": {challenge},
	}

	return AUTHORIZE_URL + params.Encode(), nil
} 

func getUserProfile(token string) (*UserProfile, error) {

	endpointUrl := API_URL + "/v1/me"

	request, err := http.NewRequest("GET", endpointUrl, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer " + token)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	var userProfileResponse UserProfile

	if err := json.NewDecoder(response.Body).Decode(&userProfileResponse); err != nil {
		return nil, err
	}

	return &userProfileResponse, nil

}