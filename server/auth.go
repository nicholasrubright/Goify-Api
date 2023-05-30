package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goify-api/client"
	"github.com/goify-api/client/models"
)

var (
	CLIENT_ID = ""
	CLIENT_REDIRECT = ""
	CLIENT_SECRET = ""
)

const (
	TOKEN_HEADER = "X-Goifiy-Token"
)

func GetAuthorizationUrl(c *gin.Context) {
	authorizationResponse, err := client.GetAuthorizationUrl(CLIENT_ID, CLIENT_REDIRECT)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, authorizationResponse)
}

func GetAccessToken(c *gin.Context) {

	var clientTokenRequest models.ClientTokenRequest

	if err := c.BindJSON(&clientTokenRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	clientResponse, clientErrorResponse := client.GetAccessToken(CLIENT_ID, CLIENT_SECRET, clientTokenRequest.Code, CLIENT_REDIRECT)

	if clientErrorResponse != nil {
		c.IndentedJSON(http.StatusInternalServerError, clientErrorResponse)
		return
	} 

	c.IndentedJSON(http.StatusFound, clientResponse)

}