package main

import (
	"encoding/json"
	"errors"
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
func getAccessToken(clientId string, code string, verifier string, redirectUri string) (*ClientTokenResponse, error) {

	formData := url.Values{
		"client_id": {clientId},
		"grant_type": {"authorization_code"},
		"code": {code},
		"redirect_uri": {redirectUri},
		"code_verifier": {verifier},
	}

	request, err := http.NewRequest("POST", TOKEN_URL, strings.NewReader(formData.Encode()))

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("HTTP Request failed", err)
		return nil, err
	}

	statusOk := response.StatusCode >= 200 && response.StatusCode < 300

	if !statusOk {
		fmt.Println("Non-Ok HTTP status: ", response.StatusCode)
		return nil, errors.New("token response is wrong")
	}

	var clientTokenResponse ClientTokenResponse 

	if err := json.NewDecoder(response.Body).Decode(&clientTokenResponse); err != nil {
		return nil, err
	}

	return &clientTokenResponse, nil
}



func getAuthUrl(clientId string, redirectUrl string) (*ClientAuthUrlResponse, error) {
	
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

	authUrl := AUTHORIZE_URL + params.Encode()

	return &ClientAuthUrlResponse{
		Url: authUrl,
		Verifier: verifier,
	}, nil
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