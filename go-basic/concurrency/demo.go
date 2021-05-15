package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	// put value to channel
	c <- payload
	//fmt.Println("Step1")
	//c <- payload
	//fmt.Println("Step2")
	//c <- payload
	//fmt.Println("Step3")
}

func main() {
	ch := make(chan int)
	go routine1(ch, 1)
	go routine1(ch, 2)
	go routine1(ch, 3)
	// Recieve value from chanel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)

	//go run1()
	//go run2()

}

// func run1() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Println("Run1 something")
// 	}
// }
// func run2() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Println("Run2 something")
// 	}
// }

/*
	ch := make(chan int, 1)
	ch <- 1
	ch <- 1

	// Recieve value from chanel
	fmt.Println(<-ch)

	// This code in comment will error, becase we set channel buffer = 1


*/
