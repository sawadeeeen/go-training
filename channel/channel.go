package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go receiver(ch)
	i := 0
	for i < 1000 {
		ch <- "st"
		i++
	}
}

func receiver(ch <-chan string) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
