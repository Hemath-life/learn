package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers")

	// var myPointer *int
	// fmt.Println(myPointer) // nil

	one := 3
	myPointer := &one

	fmt.Println(one)
	fmt.Println(myPointer)
	fmt.Println(*myPointer) // reding values
	fmt.Printf("%T", myPointer)

	*myPointer = *myPointer + 2

	fmt.Println(myPointer)
	fmt.Println(*myPointer)

}
