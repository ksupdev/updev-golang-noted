package main

import "fmt"

func main() {
	fmt.Println("Hi, this is demo How to use map has dic")

	/*
		Declaration Map
		-  map[Type of key]Type of value

	*/
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("Map value %v \n", numbers)                   // Map value map[one:1 three:3 two:2]
	fmt.Printf("Get value with key = %v \n", numbers["one"]) // Get value with key = 1

	/*
		Declaration Dynamic make by make()
	*/

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"

	fmt.Printf("Map value %v \n", colors)                   // Map value map[blue:#00f green:#0f0 red:#f00]
	fmt.Printf("Get value with key = %v \n", colors["red"]) // Get value with key = #f00

	/*
		Map real life
	*/
	var courses = make(map[string]map[string]int)
	courses["android"] = map[string]int{"price": 200}
	courses["android"]["code"] = 1234

	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 100
	courses["ios"]["code"] = 233

	fmt.Printf(" %v \n", courses)                                  //map[android:map[code:1234 price:200]]
	fmt.Printf(" get ios detail %v \n", courses["ios"])            //get ios detail map[code:233 price:100]
	fmt.Printf(" get price of ios %v \n", courses["ios"]["price"]) //get price of ios 100

}
