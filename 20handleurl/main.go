package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.com"

func main() {
	fmt.Println("Hnadling URL's in Golang")
	fmt.Println(myurl)

	//parisng url
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The Type of Query are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	//constarct a url
	partsofurl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		// Path: "/bsddsn",
		// RawPath: "shdddmm",
	}

	anotherurl := partsofurl.String()
	fmt.Println("new url: ", anotherurl)

}
