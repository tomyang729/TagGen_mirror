package main

import (
	"fmt"
	"github.com/clarifai/clarifai-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// NOTE: just for testing quinn's function
	// fmt.Print("started getting hashtags\n")
	// tags := make([]string, 3)
	// tags[0] = "redbull"
	// tags[1] = "starbucks"
	// tags[2] = "coffee"
	// // tags[3] = "tree"
	// // tags[4] = "outside"
	// retrievedTags := getPxTags(tags)
	// for _, string := range retrievedTags {
	// 	fmt.Print(string + "\n")
	// }
	// fmt.Print("finished getting hashtags")

	app := gin.Default()

	//load secret file
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}


	server := &Server{
		app:    app,
	}

	server.Configure()

	// Run on 5050 port
	app.Run(":5050")

}

func getClient(CLIENT_ID string, SECRET_KEY string) *clarifai.Client {

	return clarifai.NewClient(CLIENT_ID, SECRET_KEY)
}

