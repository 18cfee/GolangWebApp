package main

import "fmt"

func main() {
	fmt.Println("working")
	go func() {
		fmt.Println("func one")
	}()

	go func() {
		fmt.Println("func two")
	}()
	fmt.Scanln()
}
