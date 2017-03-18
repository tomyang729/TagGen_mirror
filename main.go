package main

import (
        "github.com/joho/godotenv"
        "github.com/clarifai/clarifai-go"
        "fmt"
        "os"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//load secret keys file
        err := godotenv.Load()
        if err != nil {
          fmt.Print("Error loading .env file")
        }

        CLIENT_ID := os.Getenv("CLIENT_ID")
        SECRET_KEY := os.Getenv("SECRET_KEY")
        client := clarifai.NewClient(CLIENT_ID, SECRET_KEY)

func main() {
	app := gin.Default()


        //clarify LIb Example
        info, err := client.Info()
        if err != nil {
          fmt.Println(err)
        } else {
          fmt.Printf("%+v\n", info)
        }
        // Let's get some context about these images
        urls := []string{"http://www.clarifai.com/img/metro-north.jpg", "http://www.clarifai.com/img/metro-north.jpg"}
        // Give it to Clarifai to run their magic
        tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

        if err != nil {
          fmt.Println("ERROR!")
          fmt.Println(err)
        } else {
          fmt.Printf("DATA: %+v\n", tag_data) // See what we got!
        }

	// Load static resources & templates
	customTemplate := template.Must(template.New("main").ParseGlob("resources/templates/base/*.tmpl"))
	customTemplate.ParseGlob("resources/templates/*.tmpl")
	app.SetHTMLTemplate(customTemplate)
	app.Static("/static", "resources/static")

	// Homepage endpoint
	app.GET("/", showHomePage)

	// Get hashtags
	// Params: img
	app.GET("/fetch", fetchTags)

	// Run on 5050 port
	app.Run(":5050")
}
func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home",
	})
}

func fetchTags(req *gin.Context) {
	// get img from request
	// pass in to Clarifi
	// use labels to hit Instagram/Twitter endpoint to get hashtags 
	// return array of hastags
}


