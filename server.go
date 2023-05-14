package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func getAuth(c *gin.Context) {

	authUrl, err := getAuthUrl(CLIENT_ID, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	authResponse := &ClientAuthResponse{
		Url: authUrl,
	}	

	c.IndentedJSON(http.StatusCreated, authResponse)

}

func getProfile(c *gin.Context) {

	var tokenRequest ClientTokenRequest

	if err := c.BindJSON(&tokenRequest); err != nil {
		return
	}

	profile, err := getUserProfile(tokenRequest.Token)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, profile)

}

func getToken(c *gin.Context) {
	
	var tokenResponse ClientTokenResponse

	if err := c.BindJSON(&tokenResponse); err != nil {
		return
	}

	token, err := getAccessToken(CLIENT_ID, tokenResponse.Code, tokenResponse.Verifier, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	res := struct {
		Token	string	`json:"token"`
	} { Token: token }

	c.IndentedJSON(http.StatusOK, res)

}