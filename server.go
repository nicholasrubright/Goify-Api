// Server houses all the actual router methods for Gin

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	client "github.com/goify-api/models/client"
)

const (
	TOKEN_HEADER = "X-Goifiy-Token"
)


func getTokenFromSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	if value := session.Get("mytoken"); value != nil {
		token := value.(string)
		log.Println("TOKEN IS HERE!!!!!")
			c.IndentedJSON(http.StatusOK, &client.ClientTokenResponse{
			Token: token,
			ExpiresIn: 600,
		})
		return token, nil
	} else {
		fmt.Println("VALUE WAS NULL IN SESSIONSSS!!!!!!!", value)
	}

	return "", errors.New("token could not be gotten")
}

func setTokenInSession(token string, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set("mytoken", token)
	session.Save()

	return nil
}


func postTest(c *gin.Context) {
	
	if token, err := getTokenFromSession(c); err == nil {
		log.Println("TOKEN HAS BEEN FOUND IN TOKEN!!!!")
		c.IndentedJSON(http.StatusOK, &client.ClientTokenResponse{
			Token: token,
			ExpiresIn: 60000,
		})
		fmt.Println("PENISPENIS!!!!!")
		return
	} else {
		fmt.Println("I GUEST IT DOESNT WORK!!!", err)
	}


	if err := setTokenInSession("balls", c); err != nil {
		log.Println("SESSION HAS NOT BEEN SAVED!!!!!")
	}


	c.JSON(http.StatusOK, "no balls")
}

func getAuth(c *gin.Context) {


	session := sessions.Default(c)
	var test string
	v := session.Get("test")
	if v == nil {
		test = "TEST!!!!"
	} else {
		test = v.(string)
		fmt.Println("FOUND!!!!")
		fmt.Println(test)
	}

	session.Set("test", test)
	session.Save()

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
	
	if token, err := getTokenFromSession(c); err == nil {
		log.Println("TOKEN HAS BEEN FOUND IN TOKEN!!!!")
		c.IndentedJSON(http.StatusOK, &client.ClientTokenResponse{
			Token: token,
			ExpiresIn: 60000,
		})
		fmt.Println("PENISPENIS!!!!!")
		return
	} else {
		fmt.Println("I GUEST IT DOESNT WORK!!!", err)
	}

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
		if err := setTokenInSession(clientTokenResponse.Token, c); err != nil {
			log.Println("SESSION HAS NOT BEEN SAVED!!!!!")
		}
	}

	if clientTokenResponse != nil {
		c.IndentedJSON(http.StatusOK, clientTokenResponse)
	} else if spotifyAuthorizationErrorResponse != nil {
		c.IndentedJSON(http.StatusForbidden, spotifyAuthorizationErrorResponse)
	} else {
		log.Println("Error returning json from getToken")
		return
	}
}