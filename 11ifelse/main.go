package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in Golang")

	loginCount := 15
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Something fishy"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("true")
	} else {
		fmt.Println("wrong")
	}

	// if err != nil {

	// }
}
