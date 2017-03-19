package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "github.com/clarifai/clarifai-go"
    "fmt"
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

    app.GET("/fetch", server.fetchTags)
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
func (s *Server) fetchTags(c *gin.Context) {

    //urls := []string{"http://placekitten.com/200/300"}
    //tag_data, err := s.client.Tag(clarifai.TagRequest{ URLs: urls })

    //resp, err := http.Post("https://api.instagram.com/v1/tags/search?q=" + tag + "&access_token=" + token)

    //if err != nil {
    //    fmt.Println(err)
    //} else {
    //    fmt.Printf("%+v\n", resp) // See what we got!
    //}

//client := &http.Client{}

//req, err := http.NewRequest("POST", "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs", nil)

//req.Header.Add("Authorization", `Bearer ThzzqYyVARtJdLbTQDwfRpa1FOk8w6`)
//req.Header.Add("Content-Type", `application/json`)
//resp, err := client.Do(req)



// TODO:
// get img url from c and then get JSON tags from Clarifai and pass in to Daria's function

}

