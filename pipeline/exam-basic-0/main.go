package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// id,first_name,last_name,email,gender,ip_address
type CsvData struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IpAddress string
	CustType  string
}

type CustomerData struct {
	Id           string
	Name         string
	CustomerType string
}

func getConstCustomerType() map[string]string {
	return map[string]string{
		"1": "AMC",
		"2": "ASL",
		"3": "Prospect",
	}
}

func main() {
	log.Println("Say hi !")
	c := loadData()

	for data := range c {
		log.Printf("%v", data)
	}

}

func loadData() <-chan CsvData {

	chanCsv := make(chan CsvData)
	csvFilePath := "customer.csv"

	go func() {
		file, _ := os.Open(csvFilePath)

		defer file.Close()

		reader := bufio.NewReader(file)
		r := csv.NewReader(reader)

		for {
			record, err := r.Read()

			if err == io.EOF {
				fmt.Println("EOF")
				break
			}
			if err != nil {
				log.Fatal(err)
				continue
			}

			// log.Printf(" %v \n ", record[0])

			chanCsv <- CsvData{
				Id:        record[0],
				FirstName: record[1],
				LastName:  record[2],
				Email:     record[3],
				Gender:    record[4],
				IpAddress: record[5],
			}
			close(chanCsv)
		}

	}()

	return chanCsv

}

func prepareData() {

}

func tranformData() {

}

func persiteData() {

}
