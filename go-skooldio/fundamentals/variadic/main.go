package main

import "fmt"

func main() {
	//variadic("1") => 1
	variadic("1", "2", "3") // => 1,2,3
}

func variadic(s ...string) {
	// parameter is slice of string
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
