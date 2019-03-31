package main

import "fmt"

func main() {
	a := 10
	fmt.Println("aのアドレス:", &a)
	Test(&a)
}

func Test(a *int) {
	fmt.Println("aのアドレス(関数内):", *a)
}
