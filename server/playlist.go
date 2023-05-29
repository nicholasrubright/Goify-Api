package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goify-api/client"
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

func CreatePlaylist(c *gin.Context) {

}