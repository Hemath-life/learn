package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switch")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value one")
	case 2:
		fmt.Println("dice value two")
	case 3:
		fmt.Println("dice value three")
		fallthrough // excute the upcomeing cases also 
	case 4:
		fmt.Println("dice value four")
	default:
		fmt.Println("value not found")

	}

}
