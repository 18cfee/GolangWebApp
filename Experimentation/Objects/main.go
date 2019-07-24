package main

import "fmt"

type innerObject struct {
	name     string
	age      int
	ranomVar int
}

type inside interface {
	printAge()
	getObject() innerObject
}

func (i innerObject) printAge() {
	fmt.Println(i.age)
}

func (i innerObject) getObject() innerObject {
	return i
}

func (i outerObject) printAge() {
	fmt.Println(i.age, " outer object style")
}

func (i outerObject) getObject() innerObject {
	return i.innerObject
}

type outerObject struct {
	lastName string
	name     string
	innerObject
}

func testAcceptInnerObject(i inside) {
	//t := i.innerObject
	fmt.Println("printing the name ", i.getObject().name)
	fmt.Println("about to print the inner age")
	fmt.Println(i)
	i.printAge()
}

func main() {
	test := innerObject{name: "carl", age: 5}
	test.printAge()
	outerO := outerObject{"fee", "outer", test}
	fmt.Println(outerO.innerObject.name)
	outerO.printAge()
	testAcceptInnerObject(outerO)
}
