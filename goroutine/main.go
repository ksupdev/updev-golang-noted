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
		ms.Stop(true)
	}()

	// Anonymus func (2)
	go func() {
		fmt.Printf(" %v at time %v \n", "begin Anonymus func (2)", time.Since(start))

		time.Sleep(50 * time.Millisecond)
		fmt.Printf(" %v at time %v \n", "end Anonymus func (2)", time.Since(start))
		ms.Stop(false)
	}()

	ms.Start()
	fmt.Printf(" %v at time %v \n", "End ", time.Since(start))
}
