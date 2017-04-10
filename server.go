package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Server struct {
	app      *gin.Engine
	clarifai *ClarifaiClient
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

	app.POST("/getTags", server.fetchTags)
}

func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home",
	})
}

type FetchTagsRequest struct {
	Image string `json:"image"`
}

func (s *Server) fetchTags(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Print("Error1")
	}

	var request FetchTagsRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		fmt.Print("Error2")
	}

	// TODO: somehow refresh token before it expires, without waiting for a request on an expired token
	if !s.clarifai.isAccess() {
		s.clarifai.RefreshAccesToken()
	}
	clarifaiTags, err := s.clarifai.GetTags(request.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	tags, err := getPxTags(clarifaiTags)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, tags)
}
