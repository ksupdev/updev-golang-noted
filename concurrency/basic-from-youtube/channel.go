package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		// Run without goroutine, you will see output "Sheep" only
		count("Sheep")
		count("fishh")
	*/

	c := make(chan string)
	go count("Sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
