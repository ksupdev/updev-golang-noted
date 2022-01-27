package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type externalData struct {
	inputData
	// fields fetched from external service
	relatedIds []string
}

func fetchData(ic <-chan inputData) <-chan externalData {
	oc := make(chan externalData)

	go func() {
		wg := &sync.WaitGroup{}

		for input := range ic {
			log.Printf("FetchData (Start): %v \n", input.id)
			wg.Add(1)
			go fetchFromExternalService(input, oc, wg)
			log.Printf("FetchData (End): %v \n", input.id)
		}
		log.Println("FetchData (Group Wait) ----- ")
		wg.Wait()
		log.Println("FetchData (Group End Wait) ----- ")
		close(oc)
	}()

	return oc
}

func fetchFromExternalService(input inputData, output chan<- externalData, wg *sync.WaitGroup) {
	// emulates an http request
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	related := make([]string, 0)
	for i := 0; i < rand.Intn(10); i++ {
		related = append(related, uuid.New().String())
	}
	output <- externalData{input, related}
	log.Printf("fetchFromExternalService: %v \n", len(related))
	wg.Done()
}
