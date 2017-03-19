package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"

	"github.com/joho/godotenv"
)

type Photo struct {
	Views int      `json:"times_viewed"`
	Tags  []string `json:"tags"`
}
type Photos struct {
	Data []Photo `json:"photos"`
}

type TotalItems struct {
	CurrentPage int `json:"current_page"`
}

func getPxTags(tags []ClarifyTag) ([]string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
		return nil, err
	}
	token := os.Getenv("PX_CONSUMER_KEY")

	totalTags := 0
	allTags := make(map[string]*TagData)
	for _, tag := range tags {
		//make sure to change rpp back to 100
		resp, err := http.Get("https://api.500px.com/v1/photos/search?term=" + tag.Name + "&tags=true&rpp=100&consumer_key=" + token)
		if err != nil {
			fmt.Print("Unable to retrieve photos for " + tag.Name)
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
		rawIn := json.RawMessage(body)
		bytes, err := rawIn.MarshalJSON()
		// fmt.Printf("%s\n", bytes)
		if err != nil {
			fmt.Print("something went bad\n")
			return nil, err
		}
		var response Photos
		parseErr := json.Unmarshal(bytes, &response)
		// fmt.Print(response.Data)
		if parseErr != nil {
			fmt.Print("Unable to parse the px response")
			return nil, err
		}

		for _, photo := range response.Data {
			for _, tag := range photo.Tags {
				totalTags++
				var ok bool
				_, ok = allTags[tag]
				if !ok {
					allTags[tag] = new(TagData)
					allTags[tag].TagUses = 0
					allTags[tag].TotalViews = 0
				}

				allTags[tag].TotalViews = allTags[tag].TotalViews + photo.Views
				allTags[tag].TagUses++
			}
		}
	}

	// var topTags []string
	topTags := make([]string, 0, 30)
	sortedTags := sortAlgo(allTags, totalTags)
	numOfTags := sortedTags.Len()
	if numOfTags < 30 {
		for _, tag := range sortedTags {
			topTags = append(topTags, tag.Key)
		}
	} else {
		for _, tag := range sortedTags[0:30] {
			topTags = append(topTags, tag.Key)
		}
	}

	return topTags, nil
}

func sortAlgo(wordFrequencies map[string]*TagData, totalTags int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		ratioUses := float64(pl[i].Value.TagUses) / float64(totalTags)
		pl[i].Value.SuperSecretValue = ratioUses * float64(pl[i].Value.TotalViews)

		if pl[i].Value.TagUses > 10 {
			fmt.Print(pl[i].Key + " - Tag \n")
			fmt.Print(pl[i].Value.SuperSecretValue)
			fmt.Print(" - algo value \n")
			fmt.Print(pl[i].Value.TotalViews)
			fmt.Print(" - Total views \n")
			fmt.Print(pl[i].Value.TagUses)
			fmt.Print(" - Tag Uses \n \n \n")
		}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type TagData struct {
	TotalViews       int
	TagUses          int
	SuperSecretValue float64
}
type Pair struct {
	Key   string
	Value *TagData
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	return p[i].Value.SuperSecretValue < p[j].Value.SuperSecretValue
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func stringArrToJSON(s []string) []byte {
	jsonArray, _ := json.Marshal(s)
	return jsonArray
}
