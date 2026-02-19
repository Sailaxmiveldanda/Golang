package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs Class in Golang")
	///no inheritance
	// no super keyword  and parent

	sailaxmi := User{"Sailaxmi", "abc@go.dev", true, 25}
	fmt.Println(sailaxmi)
	fmt.Printf("Sai Details are: %+v\n", sailaxmi)
	fmt.Printf("Name is %v and age is %v ", sailaxmi.Name, sailaxmi.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
