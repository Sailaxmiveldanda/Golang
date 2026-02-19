package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops Class")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//for loop

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	//fmt.Printf("Index is %v and day is %v\n", i, day)
	// 	fmt.Printf("Index is  and day is %v\n", day)
	// }

	rvallue := 1
	for rvallue < 10 {

		if rvallue == 2 {
			goto lco
		}
		// if rvallue == 5 {
		// 	rvallue++
		// 	continue
		// }
		if rvallue == 5 {
			break
		}
		fmt.Println(rvallue)
		rvallue++
	}

lco:
	fmt.Println("jump at")
}
