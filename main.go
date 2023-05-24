package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func main() {
	
	// Initialize Config
	Init()

	router := gin.Default()


	// Set Session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))


	// Set CORS Policiy
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{CLIENT_URL},
        AllowMethods:     []string{"GET,POST"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	

	// Define Routes
	router.GET("/api/auth", getAuth)
	router.POST("/api/token", getToken)
	router.GET("/api/profile", getProfile)

	router.POST("/api/test", postTest)

	// Run Server
	router.Run("localhost:8080")
}