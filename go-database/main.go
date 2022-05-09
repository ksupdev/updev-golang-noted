package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v", time.Hour*time.Duration(5))

	//https://github.com/golang/go/issues/41114
}
