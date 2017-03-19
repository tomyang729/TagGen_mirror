package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/clarifai/clarifai-go"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const (
	clarifaiApi = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
)

type Server struct {
	app    *gin.Engine
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
   Param: imgURL

*/
func fetchTags(c *gin.Context) {
	params := c.Request.URL.Query()
	image := params.Get("image")
	if image == "" {
		c.JSON(http.StatusBadRequest, "image parameter was not included")
		return
	}

	req, err := getRequestBody(image)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, "Success!")
}


/*
    Clarifai API request object body type:

    {
        "inputs": [
           {
                "data": {
                    "image": {
                        "url": "image-url"
                    }
                }
            }
        ]
    }

*/
func getRequestBody(input string) (*http.Request, error) {
	// for now, figure out how to make it one struct
	type Url struct {
		Url string `json:"url"`
	}
	type Image struct {
		Image Url `json:"image"`
	}
	type Data struct {
		Data Image `json:"data"`
	}
	type RequestBody struct {
		Inputs []Data `json:"inputs"`
	}

	// Create request body struct
	url := Url{input}
	image := Image{url}
	data := Data{image}
	inputs := make([]Data, 0, 1)
	inputs = append(inputs, data)
	reqBody := RequestBody{inputs}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("POST", clarifaiApi, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", `Bearer ThzzqYyVARtJdLbTQDwfRpa1FOk8w6`)
	req.Header.Add("Content-Type", `application/json`)
	return req, nil
}
