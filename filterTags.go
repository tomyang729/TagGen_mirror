package main

import (
	"encoding/json"
	"fmt"
)

type ClarifyAPI struct {
	Outputs []ClarifaiData `json:"outputs"`
}

type ClarifaiData struct {
	Data TagsObj `json:"data"`
}

type ClarifyTag struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	AppID string  `json:"app_id"`
	Value float64 `json:"value"`
}

type TagsObj struct {
	Tags []ClarifyTag `json:"concepts"`
}

func main() {
	str := `{  
   "status":{  
      "code":10000,
      "description":"Ok"
   },
   "outputs":[  
      {  
         "id":"bdb7163299ab49b186fa1013315d13d9",
         "status":{  
            "code":10000,
            "description":"Ok"
         },
         "created_at":"2017-03-19T00:33:05.035025Z",
         "model":{  
            "name":"general-v1.3",
            "id":"aaa03c23b3724a16a56b629203edc62c",
            "created_at":"2016-03-09T17:11:39.608845Z",
            "app_id":null,
            "output_info":{  
               "message":"Show output_info with: GET /models/{model_id}/output_info",
               "type":"concept",
               "type_ext":"concept"
            },
            "model_version":{  
               "id":"aa9ca48295b37401f8af92ad1af0d91d",
               "created_at":"2016-07-13T01:19:12.147644Z",
               "status":{  
                  "code":21100,
                  "description":"Model trained successfully"
               }
            }
         },
         "input":{  
            "id":"b64ca9387c444f82a8e2b754e8258741",
            "data":{  
               "image":{  
                  "url":"http://traveltherapytours.com/wp-content/uploads/2015/07/horse_riding_3.jpg"
               }
            }
         },
         "data":{  
            "concepts":[  
               {  
                  "id":"ai_B5f8b4Xx",
                  "name":"horse",
                  "app_id":null,
                  "value":0.99102366
               },
               {  
                  "id":"ai_Q7mGXSw0",
                  "name":"equestrian",
                  "app_id":null,
                  "value":0.9556599
               },
               {  
                  "id":"ai_N6BnC4br",
                  "name":"mammal",
                  "app_id":null,
                  "value":0.9241407
               },
               {  
                  "id":"ai_QN2GGtJC",
                  "name":"seated",
                  "app_id":null,
                  "value":0.91525424
               },
               {  
                  "id":"ai_8Qw6PFLZ",
                  "name":"recreation",
                  "app_id":null,
                  "value":0.9118946
               },
               {  
                  "id":"ai_FwtMR9mk",
                  "name":"motion",
                  "app_id":null,
                  "value":0.9017404
               },
               {  
                  "id":"ai_tLRfRp89",
                  "name":"equine",
                  "app_id":null,
                  "value":0.89018184
               },
               {  
                  "id":"ai_P9f3vfGr",
                  "name":"cavalry",
                  "app_id":null,
                  "value":0.88244057
               },
               {  
                  "id":"ai_2Zg0k8Tv",
                  "name":"mare",
                  "app_id":null,
                  "value":0.8747796
               },
               {  
                  "id":"ai_7WNVdPhm",
                  "name":"competition",
                  "app_id":null,
                  "value":0.8617599
               },
               {  
                  "id":"ai_5WW7fH4K",
                  "name":"race",
                  "app_id":null,
                  "value":0.8589672
               },
               {  
                  "id":"ai_l8TKp2h5",
                  "name":"people",
                  "app_id":null,
                  "value":0.85322106
               },
               {  
                  "id":"ai_3P5RRqwD",
                  "name":"rider",
                  "app_id":null,
                  "value":0.81064713
               },
               {  
                  "id":"ai_94cQBGlt",
                  "name":"two",
                  "app_id":null,
                  "value":0.8038043
               },
               {  
                  "id":"ai_rm7RsH7F",
                  "name":"jockey",
                  "app_id":null,
                  "value":0.7999816
               },
               {  
                  "id":"ai_V1FjkFXr",
                  "name":"leisure",
                  "app_id":null,
                  "value":0.7992622
               },
               {  
                  "id":"ai_J1hw9SFJ",
                  "name":"beach",
                  "app_id":null,
                  "value":0.79109657
               },
               {  
                  "id":"ai_mPGx9LMN",
                  "name":"action energy",
                  "app_id":null,
                  "value":0.7908335
               },
               {  
                  "id":"ai_971KsJkn",
                  "name":"track",
                  "app_id":null,
                  "value":0.785151
               },
               {  
                  "id":"ai_VRmbGVWh",
                  "name":"travel",
                  "app_id":null,
                  "value":0.7548101
               }
            ]
         }
      }
   ]
}
	`
	var tagsArr ClarifyAPI

	rawIn := json.RawMessage(str)
	bytes, err := rawIn.MarshalJSON()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &tagsArr)

	if err != nil {
		panic(err)
	}

	tagsArray := tagsArr.Outputs[0].Data.Tags
	fmt.Printf("%+v", tagsArray)

	for _, tag := range tagsArray {
		fmt.Println(tag.Name)
	}

	//WORKS=>>>>>>
	// obj := make(map[string]interface{})

	// rawIn := json.RawMessage(str)

	// bytes, err := rawIn.MarshalJSON()
	// if err != nil {
	// 	panic(err)
	// }

	// err = json.Unmarshal(bytes, &obj)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%v\n", obj["outputs"])
	////dif

	// in := `{
	//               "id":"ai_FwtMR9mk",
	//               "name":"motion",
	//               "app_id":null,
	//               "value":0.9017404
	//            }`

	// var tag ClarifyTag
	// // bytes, err := json.Marshal(in)

	// rawIn := json.RawMessage(in)
	// bytes, err := rawIn.MarshalJSON()

	// if err != nil {
	// 	panic(err)
	// }

	// err = json.Unmarshal(bytes, &tag)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", tag)
}
