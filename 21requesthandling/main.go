package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myrul = "http://google.com/get"

	response, err := http.Get(myrul)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	//fmt.Println(content) // byte code
	//fmt.Println(string(content))

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println("Meaasge is", responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8080/post"

	//fake json paylod
	requestBody := strings.NewReader(`
		{
			"CourseName" : "lajde",
			"price" : 0,
			"platform" : "youtibe.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
