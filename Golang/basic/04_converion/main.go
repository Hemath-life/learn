package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	// comma ok // error ok
	input, _ := reader.ReadString('\n') // i want to read till new line

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating ", numRating+1)
		fmt.Printf("Thanks for rating is %T", numRating)
	}
}
