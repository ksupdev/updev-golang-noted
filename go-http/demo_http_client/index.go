package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hi this is demo")
	resp, err := http.Get("http://example.com")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // Use for read data from body to byte array
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println(string(body)) // convert byte array to string
}
