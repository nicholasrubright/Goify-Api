// Server houses all the actual router methods for Gin

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	client "github.com/goify-api/models/client"
)

const (
	TOKEN_HEADER = "X-Goifiy-Token"
)

func getAuth(c *gin.Context) {

	authResponse, err := getAuthUrl(CLIENT_ID, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, authResponse)
}

func getProfile(c *gin.Context) {

	token := c.Request.Header[TOKEN_HEADER][0]

	spotifyProfileResponse, spotifyErrorResponse, err := getUserProfile(token)

	if err != nil {
		return
	}

	if spotifyProfileResponse != nil {
		c.IndentedJSON(http.StatusOK, spotifyProfileResponse)
		return
	} else if spotifyErrorResponse != nil {
		c.IndentedJSON(spotifyErrorResponse.Error.Status, spotifyErrorResponse)
		return
	} else {
		return
	}
	

}

func getToken(c *gin.Context) {
	
	var clientTokenRequest client.ClientTokenRequest

	if err := c.BindJSON(&clientTokenRequest); err != nil {
		log.Println("Error getting json from client token request")
		log.Println(err)
		return
	}

	clientTokenResponse, spotifyAuthorizationErrorResponse, err := getAccessToken(CLIENT_ID, clientTokenRequest.Code, CLIENT_REDIRECT)

	if err != nil {
		log.Println("Error when getting client token response")
		log.Println(err)
		return
	}

	if clientTokenResponse != nil {
		c.IndentedJSON(http.StatusOK, clientTokenResponse)
		return;
	} else if spotifyAuthorizationErrorResponse != nil {
		c.IndentedJSON(http.StatusForbidden, spotifyAuthorizationErrorResponse)
		return
	} else {
		log.Println("Error returning json from getToken")
		return
	}
}


func getPlaylists(c *gin.Context) {
	
	token := c.Request.Header[TOKEN_HEADER][0]

	clientPlaylistResponse, spotifyErrorResponse, err := getUserPlaylists(token)

	if err != nil {
		return
	}

	if clientPlaylistResponse != nil {
		c.IndentedJSON(http.StatusOK, clientPlaylistResponse)
		return
	} else if spotifyErrorResponse != nil {
		c.IndentedJSON(spotifyErrorResponse.Error.Status, spotifyErrorResponse)
		return
	} else {
		return
	}

}

func buildPlaylist(c *gin.Context) {
	token := c.Request.Header[TOKEN_HEADER][0]


	var clientCreatePlaylistResponse client.ClientCreatePlaylistResponse

	if err := c.BindJSON(&clientCreatePlaylistResponse); err != nil {
		log.Println("ERROR GETTING JSON FROM CLIENT CREATE PLAYLIST RESPONSE")
		log.Println(err)
		return
	}

	log.Println("Requests: ", clientCreatePlaylistResponse)

	temp, err := createPlaylist(clientCreatePlaylistResponse.User, clientCreatePlaylistResponse.Name, clientCreatePlaylistResponse.Description, token)

	fmt.Println(temp, err)

}