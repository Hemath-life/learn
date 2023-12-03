package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to struct")
	im := User{"HEM", "HEM@GAMIL.COM", "GOOD", 50} // {HEM HEM@GAMIL.COM GOOD 50}

	im.GetStatus()
	im.NewEmail() // sends the copay of the object ,here comes points
	fmt.Println(im)

}

type User struct { // capital is imported
	Name   string
	Email  string
	Status string
	Age    int
}

func (u User) GetStatus() {
	fmt.Println(u)
	fmt.Printf("user details are : %+v\n", u)
	fmt.Printf("user details are %v : %v\n", u.Name, u.Email)
}

func (u User) NewEmail() {
	u.Email = "@gmail.com" // wont affect in original value its the copy of the object
	fmt.Printf("user email is %v\n", u.Email)
}
io