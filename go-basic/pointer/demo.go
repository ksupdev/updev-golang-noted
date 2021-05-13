package pointer

import (
	"fmt"
	"strconv"
)

var count int = 0

func Demo() {
	fmt.Println("---- Hi, this is Pointer demo ----")

	/*
		Pointer is
	*/

	msg := "some message"
	fmt.Printf(" %v \n", msg)

	/*
		var msgPointer *string = msg => It's error ,becase pointer can not store value
	*/
	var msgPointer *string = &msg
	/*
		"&" use for get pointer address
		"*" use for defind pointer type
	*/

	fmt.Printf(" Value %v has pointer = %v \n", msg, msgPointer)
	// Value some message has pointer = 0xc0000962f0

	// Pass pointer value to changeMessage func
	changeMessage(msgPointer)
	fmt.Printf(" First After change value %v has pointer = %v \n", msg, msgPointer)
	//First After change value new Message 1 has pointer = 0xc000010310

	/*
		Can pass pointer value by changeMessage(&msg)

	*/
	changeMessage(&msg)
	fmt.Printf(" Second After change value %v has pointer = %v \n", msg, msgPointer)
	//Second After change value new Message 2 has pointer = 0xc000010310

}

func changeMessage(aPointer *string) {
	count++
	value := strconv.Itoa(count)
	*aPointer = "new Message " + value
}
