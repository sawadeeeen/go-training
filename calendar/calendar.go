package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s, t string
	var answer int

	sc.Scan()
	s = sc.Text()
	sc.Scan()
	t = sc.Text()

	if s == "S" {
		answer, _ = strconv.Atoi(t)
		if answer <= 1 || answer >= 64 {
			fmt.Println("不正な値です")
			os.Exit(1)
		}
		answer = answer + 1925
	}
	if s == "H" {
		answer, _ = strconv.Atoi(t)
		if answer <= 1 || answer >= 31 {
			fmt.Println("不正な値です")
			os.Exit(1)
		}
		answer = answer + 1988
	}

	fmt.Println(answer)
}
