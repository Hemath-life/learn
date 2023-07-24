package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	// comma ok // error ok
	input, _ := reader.ReadString('\n') // i want to read till new line

	// input, err := reader.ReadString('\n') // i want to read till new line

	fmt.Println("Thanks for rating", input)
	fmt.Printf("Thanks for rating is %T", input)

}
