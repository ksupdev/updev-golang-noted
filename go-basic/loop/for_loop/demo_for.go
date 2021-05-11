package for_loop

import "fmt"

func DemoLoop() {
	fmt.Println("Hi. this is Demo loop")
	fnFor()

	fnWhile()
}

func fnFor() {
	for index := 0; index < 5; index++ {
		fmt.Printf("For Index Simple print index %d \n ", index)
	}

}

func fnWhile() {
	index := 1
	for index < 5 {
		index++
		fmt.Printf("While Index Simple print index %d \n ", index)

	}
}
