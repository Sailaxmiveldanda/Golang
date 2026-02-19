package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files in Golang")
	content := "This need to go in file"

	file, err := os.Create("./mylocalfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)
	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylocalfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilerr(err)

	fmt.Println("Text data in file\n", string(databyte))
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
