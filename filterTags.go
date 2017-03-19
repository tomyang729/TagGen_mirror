package main

import (
	"encoding/json"
	"fmt"
)

// type ClarifyAPI struct {
// }

type ClarifyTag struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	AppID string  `json:"app_id"`
	Value float64 `json:"value"`
}

type TagsObj struct {
	Tags []ClarifyTag `json:"outputs"`
}

func main() {
	str := `{
	            "outputs":[
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
	`

	var tagsArr TagsObj

	rawIn := json.RawMessage(str)
	bytes, err := rawIn.MarshalJSON()

	if err != nil {
		panic(err)
	}
	// tags := make([]ClarifyTag, 0, 100)
	err = json.Unmarshal(bytes, &tagsArr)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", tagsArr)

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
