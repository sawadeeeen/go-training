package main

import "fmt"

var testVar int = showAndReturn("Declare", 1)

func init() {
	testVar = showAndReturn("Init", 2)
}

func main() {
	testVar = showAndReturn("Main", 3)
}

func showAndReturn(timing string, i int) int {
	fmt.Println(timing, ":", i)
	return i
}
