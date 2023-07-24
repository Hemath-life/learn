package main

import "fmt"

// no var style
// jwt := 300000 // error this is allowed inside the any methods

var jwt = 499 // it will be allowed

const LoginToken string = "jwt.dfkssdklfjsklf"

func main() {
	var username string = "hemath"
	fmt.Println(username)
	fmt.Printf("Variable is of type is: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type is: %T \n", isLogged)

	var smallVal1 uint8 = 255
	fmt.Println(smallVal1)
	fmt.Printf("Variable is of type is: %T \n", smallVal1)

	// var smallVal2 uint8 = 256
	// fmt.Println(smallVal2)
	// fmt.Printf("Variable is of type is: %T \n", smallVal2)

	// use this
	var smallVal3 int = 256
	fmt.Println(smallVal3)
	fmt.Printf("Variable is of type is: %T \n", smallVal3)

	// var smallVal5 float32 = 25.64564564565464
	// fmt.Println(smallVal5)
	// fmt.Printf("Variable is of type is: %T \n", smallVal5)

	var defaultVal1 int
	fmt.Println(defaultVal1)
	fmt.Printf("Variable is of type is: %T \n", defaultVal1)

	var defaultVal2 bool
	fmt.Println(defaultVal2)
	fmt.Printf("Variable is of type is: %T \n", defaultVal2)

	var defaultVal3 string
	fmt.Println(defaultVal3)
	fmt.Printf("Variable is of type is: %T \n", defaultVal3)

	// implicit type

	var hereString = "hello world"
	fmt.Println(hereString)
	fmt.Printf("Variable is of type is: %T \n", hereString)
	// hereString = 2 // error

	var hereInt = 243
	fmt.Println(hereInt)
	fmt.Printf("Variable is of type is: %T \n", hereInt)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type is: %T \n", numberOfUsers)

	fmt.Println(jwt)
	fmt.Printf("Variable is of type is: %T \n", jwt)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type is: %T \n", LoginToken)
}
