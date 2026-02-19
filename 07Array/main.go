package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Array")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Peach"

	fmt.Println("Fruits are: ", fruits)
	fmt.Println("Length is ", len(fruits))

	var veggies = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggies List is ", veggies)
	fmt.Println("Length is ", len(veggies))
}
