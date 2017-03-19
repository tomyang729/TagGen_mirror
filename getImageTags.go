package main


import (
	"encoding/json"
)

type ClarifaiObj struct {
	Outputs []ClarifaiData `json:"outputs"`
}

type ClarifaiData struct {
	Data TagsObj `json:"data"`
}

type ClarifyTag struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type TagsObj struct {
	Tags []ClarifyTag `json:"concepts"`
}

func getImageTagsArray(jsonStr string) ([]ClarifyTag, error) {
	var tagsData ClarifaiObj
	rawIn := json.RawMessage(jsonStr)

	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &tagsData)
	if err != nil {
		return nil, err
	}

	tagsArray := tagsData.Outputs[0].Data.Tags

	return filterImageTags(tagsArray), nil
}

func filterImageTags(tags []ClarifyTag) []ClarifyTag {
	res := []ClarifyTag{}

	for _, tag := range tags {
		if tag.Value >= 0.72 {
			res = append(res, tag)
		}
	}

	return res
}
