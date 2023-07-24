package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "http://lco.dev"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("response type %T", resp) // caller's responsibility to close the connection it will bery important

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(bytes)

	fmt.Println(content)
}
