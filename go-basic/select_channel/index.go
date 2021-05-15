package main

import (
	"fmt"
)

// A select blocks until one of its cases can run, then it executes that case.
func task(c, quit chan string) {
	for {
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() { // waiting for result from task
		for i := 0; i < 10; i++ {
			// Block until channel c has data
			fmt.Println(<-c)
		}
		// run when loop has finished
		quit <- "end"
	}()

	task(c, quit)
}
