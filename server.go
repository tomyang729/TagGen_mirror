package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "github.com/clarifai/clarifai-go"
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

    app.GET("/fetch", fetchTags)
}


func showHomePage(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Home",
    })
}

/*
   Get hashtags
   Param: img
 */
func fetchTags(c *gin.Context) {
    // get img from request
    // pass in to Clarifi
    // use labels to hit Instagram/Twitter endpoint to get hashtags
    // return array of hastags
}

