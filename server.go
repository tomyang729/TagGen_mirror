package main

import (
	"fmt"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"bytes"
	"strings"
	"github.com/gin-gonic/gin"
	"os"
	"net/url"
)

const (
	clarifaiApi = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
	clarifyAuth = "https://api.clarifai.com/v1/token"
)

type Server struct {
	app    *gin.Engine
}

type AccessToken struct {
	Token string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
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

	app.GET("/getTags", fetchTags)
	app.POST("/getTags", fetchTagsForPost)
}

func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Home",
	})
}

type FetchTagsRequest struct {
	Image string `json:"image"`
}
func fetchTagsForPost(c *gin.Context) {



	body, success := ioutil.ReadAll(c.Request.Body)
	if success != nil {}
	var request FetchTagsRequest
	success = json.Unmarshal(body, &request)

	req, err := getRequestBody(request.Image)

	if err != nil {
		fmt.Print("unable to get request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("failed accessing client")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		fmt.Print("bad response code from the api: ")
		fmt.Print(resp.StatusCode)
		c.JSON(http.StatusBadRequest, "Error getting response from clarifai API")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseString := string(bodyBytes)

	imageTags, err := getImageTagsArray(responseString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	tags, err := getPxTags(imageTags)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, tags)
}

/*
   Get hashtags
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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, "Error getting response from clarifai API")
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseString := string(bodyBytes)

	imageTags, err := getImageTagsArray(responseString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	tags, err := getPxTags(imageTags)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func getRequestBody(input string) (*http.Request, error) {

	clarifai_token := getAccesToken()
	// for now, figure out how to make it one struct
	type Base struct {
		Value string `json:"base64"`
	}
	type Image struct {
		Image Base `json:"image"`
	}
	type Data struct {
		Data Image `json:"data"`
	}
	type RequestBody struct {
		Inputs []Data `json:"inputs"`
	}

	// Create request body struct
	base := Base{strings.SplitN(input, ",", 2)[1]}
	image := Image{base}
	data := Data{image}
	inputs := make([]Data, 0, 1)
	inputs = append(inputs, data)
	reqBody := RequestBody{inputs}
	fmt.Print(reqBody)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(reqBody)
	req, err := http.NewRequest("POST", clarifaiApi, b)
	fmt.Print(req)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", `Bearer ` + clarifai_token.Token)
	req.Header.Add("Content-Type", `application/json`)
	return req, nil
}

func getAccesToken() AccessToken {
	CLARIFAI_ID := os.Getenv("CLARIFAI_CLIENT_ID")
	CLARIFAI_SECRET_KEY := os.Getenv("CLARIFAI_SECRET_KEY")

	responseData := authRequest(CLARIFAI_ID, CLARIFAI_SECRET_KEY)
	var token AccessToken
	rawIn := json.RawMessage(responseData)

	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Print("\nERROR Parsing authentication response")
		panic(err)
	}

	err = json.Unmarshal(bytes, &token)
	if err != nil {
		fmt.Print("\nERROR Parsing authentication json string")
		panic(err)
	}
	fmt.Print("\nRESULTS: ")
	fmt.Print(token.Token)
	fmt.Print(token.ExpiresIn)

	return token
}

func authRequest(CLIENT_ID string, SECRET_KEY string) []byte {

	resp, err := http.PostForm(clarifyAuth, url.Values{"grant_type": {"client_credentials"},
		"client_id": { CLIENT_ID },
		"client_secret": { SECRET_KEY } })

	fmt.Print("\nAuthentication Request constructed and sent")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseData,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("\nERROR parsing body request")
	}

	return responseData
}
