package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"net/url"
	"strconv"
)

const (
	pxApi = "https://api.500px.com/v1/photos/search/?"
	rpp = 50
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

func getPxTags(tags []ImageTag) ([]string, error) {

	token := os.Getenv("PX_CONSUMER_KEY")

	totalTags := 0
	allTags := make(map[string]*TagData)
	for _, tag := range tags {
		//make sure to change rpp back to 100
		requestUrl := constructUrl(tag, token)
		resp, err := http.Get(requestUrl)
		if err != nil {
			fmt.Print("Unable to retrieve photos for " + tag.Name)
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		rawIn := json.RawMessage(body)
		bytes, err := rawIn.MarshalJSON()
		if err != nil {
			fmt.Print("something went bad\n")
			return nil, err
		}
		var response Photos
		parseErr := json.Unmarshal(bytes, &response)
		if parseErr != nil {
			fmt.Printf("%s\n", bytes)
			fmt.Print("Unable to parse the px response for " + tag.Name + "\n")
			// fmt.Print(bytes)
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

func constructUrl(tag ImageTag, token string) string {
	params := url.Values{}
	params.Add("term", tag.Name)
	params.Add("tags", "true")
	params.Add("rpp", strconv.Itoa(rpp))
	params.Add("consumer_key", token)
	return pxApi + params.Encode()
}

func sortAlgo(wordFrequencies map[string]*TagData, totalTags int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		ratioUses := float64(pl[i].Value.TagUses) / float64(totalTags)
		pl[i].Value.SuperSecretValue = ratioUses * float64(pl[i].Value.TotalViews)
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
