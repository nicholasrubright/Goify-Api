package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	
	Init()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{CLIENT_URL},
        AllowMethods:     []string{"GET,POST"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	router.GET("/api/auth", getAuth)
	router.POST("/api/token", getToken)
	router.GET("/api/profile", getProfile)

	router.Run("localhost:8080")


}