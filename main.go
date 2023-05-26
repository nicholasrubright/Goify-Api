package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/goify-api/server"
)


func main() {
	
	// Initialize Config
	Init()

	// Set in server
	server.CLIENT_ID = CLIENT_ID
	server.CLIENT_REDIRECT = CLIENT_REDIRECT

	router := gin.Default()

	// Set CORS Policiy
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{CLIENT_URL},
        AllowMethods:     []string{"GET,POST"},
        AllowHeaders:     []string{"Origin,X-Goifiy-Token"},
        ExposeHeaders:    []string{"Content-Length,X-Goifiy-Token"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
	



	// Define Routes
	
	// Authentication & Authorization
	router.GET("/api/auth", server.GetAuthorizationUrl)
	router.POST("/api/token", server.GetAccessToken)
	
	// User
	router.GET("/api/profile", server.GetUserProfile)
	
	// Playlists
	router.GET("/api/playlists", server.GetUserPlaylists)
	router.POST("/api/playlists", server.CreatePlaylist)
	
	// Run Server
	router.Run("localhost:8080")
}