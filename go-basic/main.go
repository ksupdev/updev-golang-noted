package main

import (
	"fmt"
	"updev/gobasic/functions"
	"updev/gobasic/variable"
)

// main func is entry function of go app
func main() {
	fmt.Println("main action")

	fmt.Println("[---- Variable Demo ----]")
	variable.Demo()

	demoFunc()

}

func demoFunc() {
	fmt.Println("[---- Function Demo ----]")
	functions.RunDemoFunction()
	functions.PublicFunc()

}
