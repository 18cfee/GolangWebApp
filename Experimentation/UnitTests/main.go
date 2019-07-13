package main

import (
	"fmt"
)

func main() {
	fmt.Println("run this", Sum(5, 4))
}

func Sum(y int, x int) int {
	return y + x
}
