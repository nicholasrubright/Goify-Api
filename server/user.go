package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/goify-api/client"
)

func GetCurrentUserProfile(c *gin.Context) {

	token := c.Request.Header[TOKEN_HEADER][0]

	clientResponse, clientErrorResponse := client.GetCurrentUserProfile(token)

	if clientErrorResponse != nil {
		c.IndentedJSON(http.StatusInternalServerError, clientErrorResponse)
		return
	} 

	c.IndentedJSON(http.StatusFound, clientResponse)
}