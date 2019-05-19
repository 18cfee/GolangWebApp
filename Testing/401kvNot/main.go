package main

import "fmt"

var tax float64
var penalty float64
var years int
var interest float64
var AmountCont float64

func main() {
	tax = .24
	penalty = .1
	years = 40
	interest = .1
	AmountCont = 5000.0
	interest401()
	roth401k()
	standard()
}

func interest401() {
	var sum float64
	for i := 0; i < years*12; i++ {
		sum += AmountCont / 12
		sum *= (1 + interest/12)
	}
	fmt.Println("the sum here ", sum)
	sum = sum * (1 - tax) * (1 - penalty)
	fmt.Println("the sum for taking out of 401k", sum)
}

func roth401k() {
	var sum float64
	for i := 0; i < years*12; i++ {
		amountToAdd := AmountCont / 12 * (1 - tax)
		sum += amountToAdd
		sum *= (1 + interest/12)
	}
	fmt.Println("the sum here ", sum)
	sum = sum * (1 - penalty)
	fmt.Println("the sum for taking out Roth 401k", sum)
}

func standard() {
	var sum float64
	nontaxable := 0.0
	for i := 0; i < years*12; i++ {
		amountToAdd := AmountCont / 12 * (1 - tax)
		sum += amountToAdd
		nontaxable += amountToAdd
		sum *= (1 + interest/12)
	}
	fmt.Println("the sum here", sum)
	sum -= ((sum - nontaxable) * tax)
	fmt.Println("the sum for not using 401k", sum)
}
