package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		// Run without goroutine, you will see output "Sheep" only
		count("Sheep")
		count("fishh")
	*/

	/*
		// Run with goroutine, you will see output "Sheep and fish"
		go count("Sheep")
		count("fishh")
	*/

	/*
		// Run with goroutine, you do not see any output so what has happend well, because the func main goroutibe finishes before other goroutine so the program will exit
		go count("Sheep")
		go count("fishh")

		// Use for waiting input command from terminal for fix this case
		fmt.Scanln()
	*/

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("Sheep")
		wg.Done()
	}()

	go func() {
		count("fishh")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
