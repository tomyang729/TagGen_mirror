package main

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

func getImageTagsArray(jsonStr string) {
	var tagsData ClarifaiObj

	rawIn := json.RawMessage(jsonStr)
	bytes, err := rawIn.MarshalJSON()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &tagsData)

	if err != nil {
		panic(err)
	}

	tagsArray := tagsData.Outputs[0].Data.Tags

	//debugging
	// for _, tag := range tagsArray {
	// 	fmt.Println(tag.Name)
	// }
	filterImageTags(tagsArray)
}

func filterImageTags(tags []ClarifyTag) []ClarifyTag {
	res := []ClarifyTag{}

	for _, tag := range tags {
		if tag.Value >= 0.72 {
			res = append(res, tag)
			// fmt.Println(tag.Name)
		}
	}

	return res
}