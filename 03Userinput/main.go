package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome User"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading for our pizza: ")

	// comma oka syntax || , error  syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

}
