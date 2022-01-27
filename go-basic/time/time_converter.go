package main

import (
	"fmt"
	"time"
)

const (
	BasicFormat   = "2006-01-02T15:04:05Z07:00"
	BasicFormat_2 = "02/01/2006"
	AdvFormat     = "2006-01-02 3:4:5 pm"
	AdvFormat_2   = "2006-01-02 15:4:5"
	AdvFormat_3   = "2006-01-02 15:04:05"
	AdvFormat_4   = "20060102150405"
)

func main() {
	printCurrentTimeWithFormat()

	input := "Thu, 05/19/11, 10:47PM"
	layout := "Mon, 01/02/06, 03:04PM"
	convertStringToTime(input, layout) //2011-05-19 22:47:00 +0000 UTC

	convertStringToTime("310120222121", "020120061504") //2022-01-31 21:21:00 +0000 UTC

	output := convertToTime("2022-11-30", "2006-01-02").Format("02/01/2006")
	fmt.Println(output)

}

func printCurrentTimeWithFormat() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format(BasicFormat))
	fmt.Println(currentTime.Format(BasicFormat_2))
	fmt.Println(currentTime.Format(AdvFormat))
	fmt.Println(currentTime.Format(AdvFormat_2))
	fmt.Println(currentTime.Format(AdvFormat_3))
	fmt.Println(currentTime.Format(AdvFormat_4))
	/*
		2022-01-20T21:15:05+07:00
		20/01/2022
		2022-01-20 9:15:5 pm
		2022-01-20 21:15:5
		2022-01-20 21:15:05
		20220120211505
	*/
}

func convertStringToTime(input, layout string) {
	t, _ := time.Parse(layout, input)
	fmt.Println("-- output of convertStringToTime()")
	fmt.Println(t)
}

func convertToTime(input, layout string) time.Time {
	t, _ := time.Parse(layout, input)
	return t
}
