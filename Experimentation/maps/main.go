package main

import "fmt"

var mapO = map[string]string{}

func main() {
	fmt.Println("testing")
	mapO["three"] = "threeV"
	val, suc := mapO["three"]
	if suc == false {
		fmt.Println("error1")
	}
	fmt.Println(val)
}
