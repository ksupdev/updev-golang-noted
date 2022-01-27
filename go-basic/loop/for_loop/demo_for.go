package for_loop

import "fmt"

func DemoLoop() {
	fmt.Println("Hi. this is Demo loop")
	fnFor()

	fnWhile()

	fnWhileUsingBreak()
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

func fnWhileUsingBreak() {
	index := 1
	for {

		if index > 5 {
			break
		}
		fmt.Printf("While Index Useing Break Simple print index %d \n ", index)
		index++

	}
}
