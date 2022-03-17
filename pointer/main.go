package main

import "fmt"

func main() {
	d := 2
	double(&d)

	fmt.Println(d)
	// output = 4
}

func double(d *int) {
	*d = *d * 2
}
