package main

import (
	"fmt"
)

func main() { // entry point of the function

	fmt.Println("Welcome to function")
	greeting1()

	// func greeter2(){  // function inside function not allowed
	// fmt.Println("another method")
	// }

	greeter2()

	// result := adder(1, 3)

	result, message := proAdder(1, 2, 3, 4, 5, 5, 6, 6, 6, 7)
	fmt.Println(result)
	fmt.Println(message)

}
func greeting1() {
	fmt.Println("greeting1")
}
func greeter2() { // function inside function not allowed
	fmt.Println("another method")
}

// func adder(a int, b int) int { // input tpe and return type of the values
// 	return a + b
// }

func proAdder(a ...int) (int, string) {
	total := 0
	for _, val := range a {
		total += val
	}

	return total, "return from proAdder function"

}
