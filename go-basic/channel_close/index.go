package main

import (
	"fmt"
)

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}

	close(c)
}

func main() {
	ch := make(chan int, 1)
	go routine1(ch, 10)

	for {
		/*
			if you call close(..) then chanel will return default value (value=0) and channel status (ok=false)
		*/
		value, ok := <-ch
		if !ok {
			fmt.Printf("No more data %v %v \n", value, ok)
			break
		}
		fmt.Println(value)
	}

	/*
		// Loop until channel is close
		for value := range ch {
			fmt.Println(value)
		}
		fmt.Println("No more data")
	*/
}
