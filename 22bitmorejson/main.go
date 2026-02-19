package main

import (
	"encoding/json"
	"fmt"
)

// struucts
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json More")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React", 299, "google.com", "sai123", []string{"web-dev", "js"}},
		{"Mern", 199, "google.com", "rak123", []string{"full-stack", "js"}},
		{"Angular", 299, "google.com", "vas123", nil},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		 {
                "coursename": "Mern",
                "Price": 199,
                "website": "google.com",
                "tags": [ "full-stack", "js"]
        }
	`)

	var lcocourse course

	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// somecases where you want to add key value pair

	var myonlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for k, v := range myonlinedata {
		fmt.Printf("Key is  %v and Value is %v and type is %T\n", k, v, v)
	}
}
