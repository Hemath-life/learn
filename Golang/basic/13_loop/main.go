package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to loops")
	days := []string{"monday", "sunday", "thursdays", "saturday", "friday"}

	// fmt.Println(days) [monday sunday thursdays saturday friday]

	// // type one
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// // type two
	// for i := range days { // it gives index
	// 	fmt.Println(days[i])
	// }

	// // type three
	// for _, day := range days { // kind of foreach loop
	// 	fmt.Println(day)
	// }

	// // type three
	// for _, day := range days { // kind of foreach loop
	// 	fmt.Println(day)
	// }

	startValue := 1

	for startValue < len(days) {

		if startValue == 1 {
			goto err
		}
		if days[startValue] == "sunday" {
			break
		}
		if days[startValue] == "sunday" {
			continue
		}
		fmt.Println(days[startValue])

		startValue++

	}

err:
	"goto statement running "
	// label to use on goto statement
}
