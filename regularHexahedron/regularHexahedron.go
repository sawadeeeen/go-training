package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	sc.Scan()
	text := sc.Text()
	side, _ := strconv.Atoi(text)

	fmt.Println(surfaceArea(side))
}

func surfaceArea(side int) int {
	return 6 * side * side
}
