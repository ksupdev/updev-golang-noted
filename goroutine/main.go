package main

import (
	"fmt"
	"time"
	"updev-goroutine/service"
)

func main() {
	TestAdvGoRoutine()

}

func TestAdvGoRoutine() {
	start := time.Now()
	ms := service.NewMainService(start)

	fmt.Printf(" Start %v \n", start)

	// Anonymus func (1)
	go func() {
		fmt.Printf(" %v at time %v \n", "begin Anonymus func (1)", time.Since(start))
		time.Sleep(100 * time.Millisecond)
		fmt.Printf(" %v at time %v \n", "end Anonymus func (1)", time.Since(start))
		//ms.ExitChannel <- true
		ms.Stop()
	}()

	// Anonymus func (2)
	go func() {
		fmt.Printf(" %v at time %v \n", "begin Anonymus func (2)", time.Since(start))

		time.Sleep(50 * time.Millisecond)
		fmt.Printf(" %v at time %v \n", "end Anonymus func (2)", time.Since(start))
		//ms.ExitChannel <- false
		ms.Stop()
	}()

	ms.Start()
	fmt.Printf(" %v at time %v \n", "End ", time.Since(start))
}
