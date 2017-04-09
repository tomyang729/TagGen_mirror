package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	app := gin.Default()

	//load secret file
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	clarifai := NewClarifaiClient()
	clarifai.RefreshAccesToken()

	server := &Server{
		app:      app,
		clarifai: clarifai,
	}

	server.Configure()
	app.Run(":5050")

}
