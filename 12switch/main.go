package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Class")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	fmt.Println("Rnadom num is ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("two spot")
	case 3:
		fmt.Println("three spot")
		fallthrough
	case 4:
		fmt.Println("four spot")
	case 5:
		fmt.Println("five spot")
		fallthrough
	case 6:
		fmt.Println("six spot")
	default:
		fmt.Println("haha")
	}
}
