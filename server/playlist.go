package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goify-api/client"
	"github.com/goify-api/client/models"
)

func GetUserPlaylists(c *gin.Context) {

	token := c.Request.Header[TOKEN_HEADER][0]

	clientResponse, clientErrorResponse := client.GetCurrentUserPlaylists(token)

	if clientErrorResponse != nil {
		c.IndentedJSON(http.StatusInternalServerError, clientErrorResponse)
		return
	} 

	c.IndentedJSON(http.StatusFound, clientResponse)

}

func CombinePlaylists(c *gin.Context) {
	token := c.Request.Header[TOKEN_HEADER][0]

	var clientCombinePlaylistRequest models.ClientCombinePlaylistRequest

	if err := c.BindJSON(&clientCombinePlaylistRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}


	clientResponse, clientErrorResponse := client.CombinePlaylist(clientCombinePlaylistRequest.User_ID, clientCombinePlaylistRequest.Name, clientCombinePlaylistRequest.Description, clientCombinePlaylistRequest.Playlists, token)

	if clientErrorResponse != nil {
		c.IndentedJSON(http.StatusInternalServerError, clientErrorResponse)
		return
	} 

	c.IndentedJSON(http.StatusOK, clientResponse)
}