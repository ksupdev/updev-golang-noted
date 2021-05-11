package main

import (
	"fmt"
	"updev/gobasic/functions"
	"updev/gobasic/loop/for_loop"
	"updev/gobasic/variable"
)

// main func is entry function of go app
func main() {
	fmt.Println("main action")

	fmt.Println("[---- Variable Demo ----]")
	variable.Demo()

	fmt.Println("[---- Function Demo ----]")
	demoFunc()

	fmt.Println("[---- For loop Demo ----]")
	demoForLoop()

}

func demoFunc() {

	functions.RunDemoFunction()
	functions.PublicFunc()

}

func demoForLoop() {
	for_loop.DemoLoop()
}
