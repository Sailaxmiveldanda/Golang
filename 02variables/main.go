package main

import "fmt"

const LoginToken string = "dsfdfas" //public

func main() {
	var username string = "Sailaxmi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isverify bool = false
	fmt.Println(isverify)
	fmt.Printf("Variable is of type: %T \n", isverify)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallfloat float64 = 255.45846556
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//default values and aliases

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	//implicit type
	var website = "leetcode.com"
	fmt.Println(website)

	// no var style
	numberofusers := 1234
	fmt.Println(numberofusers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
