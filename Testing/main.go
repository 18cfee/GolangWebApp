package main

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("this function is running")
}

func main() {
	sampleString := "13"
	i, err := strconv.ParseInt(sampleString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	sampleStringTwo := []byte("testing")
	fmt.Println(sampleStringTwo)
}
