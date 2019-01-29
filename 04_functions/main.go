package main

import (
	"fmt"
)

func myFunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

// Anonymous Functions
func addOne() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println("------ Functions ------")

	fullName, err := myFunction("Ew", "Jiang")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)

	myFunc := addOne()
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())

	fmt.Println("------ Functions End ------")
}
