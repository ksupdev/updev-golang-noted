package main

import (
	"fmt"
	"time"
)

/*
go mod init updev.co.th/time
*/

func main() {

	getThaiTime()
	getJapanTime()

	getThaiTimeToString()
}

func getThaiTime() {
	locat, error := time.LoadLocation("Asia/Jakarta")

	if error != nil {
		panic(error)
	}

	currentTime := time.Now()
	fmt.Printf("Thai time : %v \n", currentTime.In(locat))
}

func getThaiTimeToString() {
	locat, error := time.LoadLocation("Asia/Jakarta")

	if error != nil {
		panic(error)
	}

	currentTime := time.Now().In(locat)
	fmt.Println(currentTime.String())
	fmt.Println(currentTime.Format("2006/01/02 15:04:05.000000"))
	fmt.Println(currentTime.Format("01:00:00.000000"))
}

func getJapanTime() {
	locat, error := time.LoadLocation("Japan")
	if error != nil {
		panic(error)
	}

	fmt.Printf("Japan time : %v \n", time.Now().In(locat))
}
