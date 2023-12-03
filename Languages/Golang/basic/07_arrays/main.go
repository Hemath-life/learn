package main

import "fmt"

func main() {

	fmt.Println("Welcome to Arrays")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "orange"

	fruits[3] = "water"

	fmt.Println(fruits)
	fmt.Println("length", len(fruits)) // show the defined length

	var veggies = [4]string{"potato", "tomato"}

	fmt.Println(veggies)
	fmt.Println("length", len(veggies)) // show the defined length

}
