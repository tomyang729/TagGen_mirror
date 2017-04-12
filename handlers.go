package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home",
	})
}

// The 'main' controller to fetch tags via clarifai and px500 clients
/* use dependency injection to pass in clarifai (and px for later) client in order to make
 * the function more testable and modular;
 * return gin.Handler to satisfy Gin's router method
 */
func fetchTags(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Printf("Error--read file: %s", err)
	}
	var request struct {
		Image string `json:"image"`
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		fmt.Printf("Error--parse json: %s", err)
	}

	// TODO: somehow refresh token before it expires, without waiting for a request on an expired token
	if !isCLFAccessible(CLFclient) {
		refreshCLFToken(CLFclient)
	}
	clarifaiTags, err := getImageTagsFromCLF(request.Image, CLFclient)
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
