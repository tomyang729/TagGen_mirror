package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

const (
	clarifaiApi = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
	clarifyAuth = "https://api.clarifai.com/v1/token"
	minValue    = 0.72
	maxTags     = 10
)

type ClarifaiClient struct {
	Token  string
	Expiry int
	Access bool
}

type ImageTag struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func NewClarifaiClient() *ClarifaiClient {
	return &ClarifaiClient{
		Access: false,
	}
}

// Check whether the client still has access rights to Clarifai
func isCLFAccessible() bool {
	if !CLFclient.Access || CLFclient.Expiry <= 0 {
		CLFclient.Access = false
		return false
	}
	return true
}

// Refresh the access token
func refreshCLFToken() {
	CLARIFAI_ID := os.Getenv("CLARIFAI_CLIENT_ID")
	CLARIFAI_SECRET_KEY := os.Getenv("CLARIFAI_SECRET_KEY")

	responseData := authRequest(CLARIFAI_ID, CLARIFAI_SECRET_KEY)
	var token struct {
		//type AccessToken struct {
		Token  string `json:"access_token"`
		Expiry int    `json:"expires_in"`
		//}
	}
	rawIn := json.RawMessage(responseData)

	b, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Println("ERROR Parsing authentication response")
		panic(err)
	}

	err = json.Unmarshal(b, &token)
	if err != nil {
		fmt.Println("ERROR Parsing authentication json string")
		panic(err)
	}

	fmt.Printf("Token: %v\n", token.Token)
	fmt.Printf("Expiry: %v\n", token.Expiry)

	CLFclient.Token = token.Token
	CLFclient.Expiry = token.Expiry
	CLFclient.Access = true

}

func authRequest(CLIENT_ID string, SECRET_KEY string) []byte {
	resp, err := http.PostForm(clarifyAuth, url.Values{"grant_type": {"client_credentials"},
		"client_id":     {CLIENT_ID},
		"client_secret": {SECRET_KEY}})

	fmt.Print("\nAuthentication Request constructed and sent")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("\nERROR parsing body request")
	}

	return responseData
}

// Get the tags from the Clarifai API
func getImageTagsFromCLF(image string) ([]ImageTag, error) {
	requestBody, err := getRequestBody(image)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", clarifaiApi, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", `Bearer `+ CLFclient.Token)
	req.Header.Add("Content-Type", `application/json`)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseString := string(bodyBytes)

	imageTags, err := getImageTagsArray(responseString)
	if err != nil {
		return nil, err
	}

	return imageTags, nil
}

// Structure the given image to a request body with Clarifai's format
func getRequestBody(input string) (*bytes.Buffer, error) {
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
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Get image tag array from Clarifai's API response
func getImageTagsArray(jsonStr string) ([]ImageTag, error) {
	type TagsObj struct {
		Tags []ImageTag `json:"concepts"`
	}

	type ClarifaiData struct {
		Data TagsObj `json:"data"`
	}

	type ClarifaiObj struct {
		Outputs []ClarifaiData `json:"outputs"`
	}

	var tagsData ClarifaiObj
	rawIn := json.RawMessage(jsonStr)

	b, err := rawIn.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &tagsData)
	if err != nil {
		return nil, err
	}

	tagsArray := tagsData.Outputs[0].Data.Tags

	return filterImageTags(tagsArray), nil
}

// clarifaiTags is a slice of ImageTag. Used for sorting / filtering
type clarifaiTags []ImageTag

func (slice clarifaiTags) Len() int {
	return len(slice)
}

func (slice clarifaiTags) Less(i, j int) bool {
	return slice[i].Value < slice[j].Value
}

func (slice clarifaiTags) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice clarifaiTags) Filter(val float64) clarifaiTags {
	filtered := make([]ImageTag, 0, slice.Len())
	for _, tag := range slice {
		if tag.Value >= val {
			filtered = append(filtered, tag)
		}
	}
	return filtered
}

// Filter tags based on minScore, sort, and trim if there are more than maxTas
func filterImageTags(tags clarifaiTags) []ImageTag {
	tags = tags.Filter(minValue)

	sort.Sort(tags)

	if tags.Len() > maxTags {
		tags = tags[:maxTags]
	}

	return tags
}
