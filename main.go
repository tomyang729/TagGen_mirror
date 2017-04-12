package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"html/template"
)

// Clarifai client and 500PX client singleton
// To keep it simple, make them global so that we don't need to pass them around
var CLFclient = NewClarifaiClient()
//var PxClient = NewPxClient()  // TODO: later, if we want to post photos to 500px

func main() {
	// Create router
	router := gin.Default()

	// TODO: Set up environment var (like prod/dev/test)

	/* Here we load secret file for now. This should be removed later
	 * The api secrets should be pre-set on the hosting machine and accessed via envVar during runtime
         */
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Load static resources & templates
	customTemplate := template.Must(template.New("main").ParseGlob("resources/templates/base/*.tmpl"))
	customTemplate.ParseGlob("resources/templates/*.tmpl")
	router.SetHTMLTemplate(customTemplate)
	router.Static("/static", "resources/static")

	// Define routes **currently we only have two; it's not necessary to separate it to another file yet
	router.GET("/", showHomePage)
	router.POST("/getTags", fetchTags)

	// Listen and serve
	router.Run(":5050")
}
