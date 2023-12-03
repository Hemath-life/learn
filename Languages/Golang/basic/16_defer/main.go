package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("world")      // this line will be executed when all code ran
	defer fmt.Println("one two ")   // this line will be executed when all code ran
	defer fmt.Println("one three ") // this line will be executed when all code ran

	// upcoming three will be print like
	// one three
	// one two
	// world
	// first in last out

	fmt.Println("hello")

}
