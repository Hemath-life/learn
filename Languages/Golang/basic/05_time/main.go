package main

import (
	"fmt"

	"time"
)

func main() {
	fmt.Println("welcome to type study")

	currentTime := time.Now()
	fmt.Println(currentTime)

	// currentTime := time.Now()

	fmt.Println(currentTime.Format("2006-01-02 15:04:05 Monday")) // format to 01-02-2006 this is format string

	// createDate := time.Date(2022, time.April, 10, 23, 23, 2, 1, time.UTC)
	// fmt.Println(createDate)

	// // currentTime := time.Now()
	// fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday")) // format to 01-02-2006 this is format string

}
