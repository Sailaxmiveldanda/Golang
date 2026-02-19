package main

import "fmt"

func main() {
	fmt.Println("Class on Functions")
	greeter()

	result := add(2, 3)
	fmt.Println("Result is:", result)
	proresult, mess := proadder(1, 2, 3, 4, 5, 6)
	fmt.Println("Pro result is ", proresult)
	fmt.Println(mess)

	//greeter2()
}

// func greeter2() {
// 	fmt.Println("Another Method")
// }

func greeter() {
	fmt.Println("Hello Sai")
}

func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi, Pro result function"
}

func add(v1 int, v2 int) int {
	return v1 + v2
}
