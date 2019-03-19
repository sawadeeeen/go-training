package main

import (
	"fmt"
)

func main() {
	num := [5]int{1, 2, 3, 4, 5}

	var slice1 []int
	var slice2 []int

	slice1 = num[:]
	fmt.Println(slice1)

	slice2 = num[1:4]
	fmt.Println(slice2)
}
