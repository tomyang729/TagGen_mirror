package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "github.com/clarifai/clarifai-go"
    "net/http"
)

type Server struct {
    app *gin.Engine
    client *clarifai.Client	 
}

func (server *Server) Configure() {
    app := server.app

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
}


func showHomePage(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Home",
    })
}

func fetchTags(c *gin.Context) {
    _ = clarifai.NewClient("a", "b") // just so it will compile
    // get img from request
    // pass in to Clarifi
    // use labels to hit Instagram/Twitter endpoint to get hashtags
    // return array of hastags
}

