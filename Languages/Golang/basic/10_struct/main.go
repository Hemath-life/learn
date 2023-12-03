package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to struct")
	hema := User{"HEM", "HEM@GAMIL.COM", "GOOD", 50} // {HEM HEM@GAMIL.COM GOOD 50}

	fmt.Println(hema) // {HEM HEM@GAMIL.COM GOOD 50}

	fmt.Printf("HEMATH details are : %+v\n", hema) // +v need to user for structure
	// HEMATH details are : {Name:HEM Email:HEM@GAMIL.COM Status:GOOD Age:50}

	fmt.Printf("HEMATH details are %v : %v\n", hema.Name, hema.Email) // HEMATH details are HEM : HEM@GAMIL.COM

}

type User struct { // capital is imported
	Name   string
	Email  string
	Status string
	Age    int
}
