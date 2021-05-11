package main

import (
	"fmt"
	"updev/gobasic/foreach_range"
	"updev/gobasic/functions"
	"updev/gobasic/if_else_switch_case"
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

	fmt.Println("[---- Foreach with Range Demo ----]")
	dempForEachWithRange()

	fmt.Println("[---- If else and Switch case Demo ----]")
	dempIfElseSwithcCase()

}

func demoFunc() {

	functions.RunDemoFunction()
	functions.PublicFunc()

}

func demoForLoop() {
	for_loop.DemoLoop()
}

func dempForEachWithRange() {
	foreach_range.DemoForeachRange()
}

func dempIfElseSwithcCase() {
	if_else_switch_case.Demo()
}
