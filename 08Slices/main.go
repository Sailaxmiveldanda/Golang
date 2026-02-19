package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on Slices")

	var fruitsList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruits is %T\n", fruitsList)

	fruitsList = append(fruitsList, "mango", "banana")
	fmt.Println(fruitsList)

	//fruitsList = append(fruitsList[1:3])
	fmt.Println(fruitsList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 867
	//highscore[4] = 777
	highscore = append(highscore, 777, 555, 123)
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	//remove value from slice based on index
	var courses = []string{"react", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
