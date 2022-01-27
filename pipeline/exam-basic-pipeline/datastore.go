package main

import (
	"log"
	"time"
)

type saveResult struct {
	idsSaved  []string
	timestamp int64
}

func saveData(ic <-chan externalData) <-chan saveResult {
	oc := make(chan saveResult)

	go func() {
		const batchSize = 7
		batch := make([]string, 0)
		for input := range ic {
			if len(batch) < batchSize {
				batch = append(batch, input.inputData.id)
				// log.Printf("saveData build (%v) ----- \n ", input.inputData.id)
				log.Println("saveData build")
			} else {
				log.Println("saveData save")
				oc <- persistBatch(batch)
				batch = []string{input.inputData.id}
				// log.Printf("saveData persistBatch (%v) ----- \n ", input.inputData.id)

			}
		}

		if len(batch) > 0 {
			log.Println("saveData build 32")
			oc <- persistBatch(batch)
		}

		close(oc)
	}()

	return oc
}

func persistBatch(batch []string) saveResult {
	return saveResult{batch, time.Now().UnixNano()}
}
