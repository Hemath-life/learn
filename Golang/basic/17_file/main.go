package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in  golang")

	content := "This needs to go in a file"

	file, err := os.Create("./myFile.text")

	if err != nil {
		panic(err) // // throwing the error stop the code
	}

	length, err := io.WriteString(file, content)

	checkNilError(err)

	defer file.Close() // execute in last
	fmt.Println(length)

	read
}

func read(filename string) {

	bytes, err = ioutil.ReadFile(filename) // data will come in byte

	checkNilError(err)

	fmt.Println(byte)
	fmt.Println(string(byte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err) // throwing the error stop the code
	}
}
