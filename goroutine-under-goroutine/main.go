package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Printf("---- Start ---- [%v] \n", startTime)
	fmt.Printf("---- Start program ---- [%v] \n", time.Since(startTime))
	inputData := "123"
	go func1(inputData, startTime)

	go func2(inputData, startTime)

	go func(s string) {
		go func3(s, startTime)
		for _, d := range s {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Anonymus-1 output = %c [%v]\n", d, time.Since(startTime))
		}
	}(inputData)

	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("---- End progra ---- [%v] \n", time.Since(startTime))
	fmt.Printf("---- End ---- [%v] \n", startTime)

}

func func1(s string, startTime time.Time) {
	for _, d := range s {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("func1 output = %c [%v]\n", d, time.Since(startTime))
	}
}

func func2(s string, startTime time.Time) {
	for _, d := range s {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("func2 output = %c [%v]\n", d, time.Since(startTime))
	}
}

func func3(s string, startTime time.Time) {
	for _, d := range s {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("func3 output = %c [%v]\n", d, time.Since(startTime))
	}
}
