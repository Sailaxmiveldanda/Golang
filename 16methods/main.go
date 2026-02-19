package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Class on Methods")

	fmt.Println("Welcome to Structs Class in Golang")
	///no inheritance
	// no super keyword  and parent

	sailaxmi := User{"Sailaxmi", "abc@go.dev", true, 25}
	fmt.Println(sailaxmi)
	fmt.Printf("Sai Details are: %+v\n", sailaxmi)
	fmt.Printf("Name is %v and age is %v \n", sailaxmi.Name, sailaxmi.Age)
	sailaxmi.GetStatus()
	sailaxmi.NewMail()
	fmt.Printf("Name is %v and age is %v \n", sailaxmi.Name, sailaxmi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int not exportable
}

func (u User) GetStatus() {
	fmt.Println(os.Stdout, "Is user Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of User is: ", u.Email)
}
