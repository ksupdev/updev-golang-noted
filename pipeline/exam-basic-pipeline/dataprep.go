package main

import (
	"log"
	"time"
)

type inputData struct {
	id        string
	timestamp int64
}

func prepareData(ic <-chan CusUUID) <-chan inputData {
	oc := make(chan inputData)
	go func() {
		for id := range ic {
			input := inputData{id: id.Id, timestamp: time.Now().UnixNano()}
			log.Printf("Data ready for processing: %v \n", id.Id)
			// log.Println(line)
			oc <- input
		}
		close(oc)
	}()

	return oc
}
