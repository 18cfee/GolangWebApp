package main

import (
	"fmt"
	"strconv"
)

func main() {
	sampleString := "13"
	i, err := strconv.ParseInt(sampleString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
