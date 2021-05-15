package main

import "fmt"

func main() {
	fmt.Println("Hi, this is demo Defer")

	defer fmt.Println("first defer")
	defer fmt.Println("seconde defer")
	defer fmt.Println("third defer")

	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}

	/*
		defer :=> it's last action before end method

		 	0
			1
			2
			3
			4
			5
			6
			7
			8
			9
			third defer
			seconde defer
			first defer
	*/
}
