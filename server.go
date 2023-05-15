package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func getAuth(c *gin.Context) {

	authResponse, err := getAuthUrl(CLIENT_ID, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, authResponse)

}

func getProfile(c *gin.Context) {

	var clientProfileRequest ClientProfileRequest

	if err := c.BindJSON(&clientProfileRequest); err != nil {
		return
	}

	profile, err := getUserProfile(clientProfileRequest.Token)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, profile)

}

func getToken(c *gin.Context) {
	
	var clientTokenRequest ClientTokenRequest

	if err := c.BindJSON(&clientTokenRequest); err != nil {
		return
	}

	clientTokenResponse, err := getAccessToken(CLIENT_ID, clientTokenRequest.Code, clientTokenRequest.Verifier, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, clientTokenResponse)
}