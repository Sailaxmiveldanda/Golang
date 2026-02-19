// invokes a function whose execution is deferred to the moment
// the surrounding funtion returns
// deffered funtions are invoked immediately before the surrounding function
// returns in the reverse orderr they deffered(LIFO)
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello ")
	mydefer()
	//defer fmt.Println("Hello World")
}

// world,one,two -> hello, two , one, world

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
