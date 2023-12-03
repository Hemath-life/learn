package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	var fruitsList = []string{"apple", "peach"}

	fmt.Printf("type of fruits list %T \n", fruitsList) // []string

	// adding values
	fruitsList = append(fruitsList, "mango")
	fmt.Println(fruitsList) //[apple peach mango]

	// deleting first values
	fruitsList = append(fruitsList[1:])
	fmt.Println(fruitsList) // [peach mango]

	// slicing
	higScores := make([]int, 4)
	higScores[0] = 12
	higScores[1] = 13
	higScores[2] = 14
	higScores[3] = 15

	// higScores[4] = 16 // error out of range
	higScores = append(higScores, 16, 5, 45, 45) // make new slice the append the values it will includes memory
	fmt.Println(higScores)                       // [12 13 14 15 16 5 45 45]

	// sorting
	sort.Ints(higScores)
	fmt.Println(higScores)                     // [5 12 13 14 15 16 45 45]
	fmt.Println(sort.IntsAreSorted(higScores)) // true

	// remove values in slice based on index
	var courses = []string{"react", "javascript", "express", "swift", "python", "ruby"}
	fmt.Println(courses) // [react javascript express swift python ruby]

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // [react javascript swift python ruby]

}
