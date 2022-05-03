package fibo

import "fmt"

func program() {

	// num := 20
	jobs := make(chan int, 20)
	results := make(chan int, 20)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	for i := 0; i < 20; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 20; j++ {
		v := <-results
		fmt.Printf("-- %v \n", v)
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	// jobs <-chan outbound (receive) receive (<- chan)
	// result chan <- inbound (send)  sends (chan <-)
	fmt.Printf("==worker : %v == \n", id)
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
