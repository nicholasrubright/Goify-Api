package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	
	// Initialize Config
	Init()

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
	router.GET("/api/auth", getAuth)
	router.POST("/api/token", getToken)
	router.GET("/api/profile", getProfile)
	router.GET("/api/playlists", getPlaylists)
	router.POST("/api/playlists", buildPlaylist)
	
	// Run Server
	router.Run("localhost:8080")
}