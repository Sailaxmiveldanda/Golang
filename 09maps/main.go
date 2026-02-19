package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["py"] = "Pyhton"

	fmt.Println("List of languages: ", languages)
	fmt.Println("Js shorts for", languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)

	//loops
	for key, value := range languages {
		fmt.Printf("For Key %v , Value is %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("For Key v, Value is %v\n", value)
	}

}
