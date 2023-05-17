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

	spotifyProfileResponse, spotifyErrorResponse, err := getUserProfile(clientProfileRequest.Token)

	if err != nil {
		return
	}

	if spotifyProfileResponse != nil {
		c.IndentedJSON(http.StatusOK, spotifyProfileResponse)
	} else if spotifyErrorResponse != nil {
		c.IndentedJSON(spotifyErrorResponse.Error.Status, spotifyErrorResponse)
	} else {
		return
	}
	

}

func getToken(c *gin.Context) {
	
	var clientTokenRequest ClientTokenRequest

	if err := c.BindJSON(&clientTokenRequest); err != nil {
		return
	}

	clientTokenResponse, spotifyAuthorizationErrorResponse, err := getAccessToken(CLIENT_ID, clientTokenRequest.Code, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	if clientTokenResponse != nil {
		c.IndentedJSON(http.StatusOK, clientTokenResponse)
	} else if spotifyAuthorizationErrorResponse != nil {
		c.IndentedJSON(http.StatusForbidden, spotifyAuthorizationErrorResponse)
	} else {
		return
	}
}