package main

import (
	"log"
)

type CusUUID struct {
	Id string
}

func main() {
	log.Println("Hi")
	c := saveData(fetchData(
		prepareData(
			generateData(),
		),
	))

	for data := range c {
		log.Printf("Items saved: %+v", len(data.idsSaved))
	}
}
