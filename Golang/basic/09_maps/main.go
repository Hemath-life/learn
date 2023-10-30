package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println(languages)       // map[go:golang js:Javascript py:python]
	fmt.Println(languages["js"]) // javascript

	delete(languages, "go")
	fmt.Println(languages) //    map[js:javascript py:python]

	// loops are interesting

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}

}
