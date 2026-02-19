package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	//pointer declaration
	// var ptr *int
	// fmt.Println("value of pointer is: ", ptr)

	//referenmce  - &
	mynumber := 23
	var ptr = &mynumber
	fmt.Println("Value of actual pointer: ", ptr)
	fmt.Println("Value of actual pointer: ", *ptr)

	//*ptr -- value

	*ptr = *ptr * 2
	fmt.Println("Value is ", mynumber)

}
