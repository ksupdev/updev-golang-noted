package main

import "fmt"

func main() {
	fmt.Println("this's a demo array")

	//array declaration and defind values
	var array1 []int = []int{1, 2, 3, 4}
	/*
		The Other ways
		- var array2 = []int{1, 2, 3, 4}
		- array3 := []int{1, 2, 3, 4}
		- array4 := [3]string
	*/

	fmt.Printf("Print value in array1 %d \n", array1)
	for _, item := range array1 {
		fmt.Printf("- %d \n", item)
	}

	var arrayString [3]string
	arrayString[0], arrayString[1], arrayString[2] = "value1", "value2", "value3"
	fmt.Printf("Print value in arrayString %v \n", arrayString)
	for _, item := range arrayString {
		fmt.Printf(" - %v \n", item)
	}

}
