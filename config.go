package main

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	CLIENT_ID = ""
	CLIENT_REDIRECT = ""
	CLIENT_URL = ""
	CLIENT_SECRET = ""
)


func Init() {
	godotenv.Load()

	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_REDIRECT = os.Getenv("CLIENT_REDIRECT")
	CLIENT_URL = os.Getenv("CLIENT_URL")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
}