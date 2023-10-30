package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to if else ")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "you can login "
	} else if loginCount == 10 {
		result = "you have last one login "
	} else {
		result = "you have reached login count "
	}

	fmt.Println(result)

	if num := 3; num < 3 {
		fmt.Println("num is less then 3 ")
	} else {
		fmt.Println("num is greater then 3 ")
	}

}
