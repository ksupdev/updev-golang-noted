package main

import "fmt"

func main() {
	fmt.Println("Hi, this's demo slice")
	/*
		Array requir to setup size but Slice is not requir (dynamic array)
	*/
	var numbers = make([]int, 3, 5) // specify a capacity (5), pass a third argument
	/*
		[]int = type
		3 = len
		5 = cap
	*/
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	showSlice(numbers)
	//len of slice 5 : [0 0 0 1 2] and cap : 5

	numbers = append(numbers, 3)
	showSlice(numbers)
	/*
		if you append value more than the capacity value in make(), golang will autoscale capacity for support this slice
		, default 5 scale up to 10
	*/
	//len of slice 6 : [0 0 0 1 2 3] and cap : 10

	// create slice has not allocate cap and len
	var numbers2 []int
	showSlice(numbers2)

	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)

	showSlice(numbers2)

	fmt.Println("------ Show Slice remove ------")
	var arrayNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	showSlice(arrayNumbers)

	// leading remove : remove first value of slice
	arrayNumbers = arrayNumbers[1:] //remove 1 from 1, 2, 3, 4, 5, 6, 7, 8, 9
	showSlice(arrayNumbers)         // =>len of slice 8 : [2 3 4 5 6 7 8 9] and cap : 8
	arrayNumbers = arrayNumbers[1:] //remove 2 from 2, 3, 4, 5, 6, 7, 8, 9
	showSlice(arrayNumbers)         // =>l en of slice 7 : [3 4 5 6 7 8 9] and cap : 7

	//trailing remove
	arrayNumbers = arrayNumbers[0 : len(arrayNumbers)-1] //remove 9 from 3 4 5 6 7 8 9
	showSlice(arrayNumbers)                              // =>len of slice 6 : [3 4 5 6 7 8] and cap : 7

	fmt.Println("------ Show Slice remove specific index ------")
	var arrayNumbersFortestRemove = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	resultValue := removeIndex(arrayNumbersFortestRemove, 3)
	showSliceString(resultValue)

	fmt.Println("----- Noted select slice ------")
	selectSllice()
}

func showSlice(x []int) {
	fmt.Printf("len of slice %v : %d and cap : %d \n", len(x), x, cap(x))
}

func showSliceString(x []string) {
	fmt.Printf("len of slice %v : %v and cap : %d \n", len(x), x, cap(x))
}

func removeIndex(s []string, index int) []string {
	fmt.Printf("s[:index] => %v \n", s[0:index])
	fmt.Printf("s[index+1:] => %v \n", s[index+1:])

	return append(s[:index], s[index+1:]...)
	/*
		if index = 3
		and s = A B C D E F G H I
		s[:index] => [A B C]
		s[index+1:] => [E F G H I]

	*/
}

func selectSllice() {
	demo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	demo = append(demo, 10)
	demo = append(demo, 11)
	fmt.Printf("Set of values %v has len %d and cap %d \n", demo, len(demo), cap(demo))
	//Set of values [1 2 3 4 5 6 7 8 9 10 11] has len 11 and cap 18

	fmt.Printf("Select demo[1:4] %v \n", demo[1:4])
	//Select demo[1:4] [2 3 4]

	fmt.Printf("Select demo[3:4] %v \n", demo[3:4])
	//Select demo[3:4] [4]

	fmt.Printf("Select demo[4:4] %v \n", demo[4:4])
	//Select demo[4:4] []

	fmt.Printf("Select demo[3:5] %v \n", demo[3:5])
	//Select demo[3:4] [4]

	fmt.Printf("Select demo[0:4] vs demo[:4] => %v vs %v \n", demo[0:4], demo[:4])
	//Select demo[0:4] vs demo[:4] => [1 2 3 4] vs [1 2 3 4]

	fmt.Printf("Select demo[2:len(demo)] vs demo[2:] => %v vs %v \n", demo[2:len(demo)], demo[2:])
	//Select demo[2:len(demo)] vs demo[:4] => [3 4 5 6 7 8 9 10 11] vs [3 4 5 6 7 8 9 10 11]
}

/*
	Slice noted
		slice_variable[index of array:number of value]
		slice_variable[3:5] of [1 2 3 4 5 6 7 8 9 10 11] ===> [4,5]
		becase
			- index of array = 3 => 1[0],2[1],3[2],4[3]  ** value[index]
			- number of value = 4 => 1,2,3,4,5 => 4

	About slice
		Remove slice is not remove value in array but it's select value new value and replace it


*/
