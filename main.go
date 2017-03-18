package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"github.com/clarifai/clarifai-go"
)

func main() {
	app := gin.Default()

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


