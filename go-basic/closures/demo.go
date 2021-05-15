package main

import "fmt"

func main() {
	fmt.Println("Hi, This is Closures demo")

	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInt := intSeq()
	fmt.Println(newInt())  // 1
	fmt.Println(nextInt()) // 4
	fmt.Println(newInt())  // 2

	seqString := stringSeq()
	fmt.Println(seqString()) // Y = 1
	fmt.Println(seqString()) // Y = 2
	fmt.Println(seqString()) // Y = 3

	// create new variabel
	fmt.Println(stringSeq()()) // Y = 1
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf(" Y = %d \n", y)
	}
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
