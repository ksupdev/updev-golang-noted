package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Response time: ", time.Since(start))
	}()

	var wg sync.WaitGroup
	var urls = []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/users/1",
		"https://jsonplaceholder.typicode.com/posts/1",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Fetch the URL.
			resp, err := http.Get(url)
			if err != nil {
				log.Panic(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}

			log.Println(url)
			log.Println(string(body))
		}(url)
	}

	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
