package main

import (
	"fmt"

	"net/url"
)

var myUrl = "http://lco.dev:3000/learn?coursename=reactjs&courseid=434847538"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	// parrsing
	res, _ := url.Parse(myUrl)

	// fmt.Println(res)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	// fmt.Println(res.Port())

	params := res.Query()

	fmt.Println(params)
	fmt.Printf("%T\n", params) // url.Values

	fmt.Println(params["courseid"]) // [434847538]

	for _, val := range params {
		fmt.Println(val)
	}

	patsOfUrl := &url.URL{ // way to construct new url
		Scheme:  "https",
		Host:    "lec.dev",
		Path:    "/learn",
		RawPath: "course=react",
	}

	newulr := patsOfUrl.String()

	fmt.Println(newulr)
}
