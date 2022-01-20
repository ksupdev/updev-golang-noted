package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "123321"
	output := replaceWith(input, "-", "/")
	fmt.Println(output) // 123321

	input = "123-321"
	output = replaceWith(input, "-", "/")
	fmt.Println(output) // 123/321

	input = "123-32-1" // 123/32-1
	output = replaceWith(input, "-", "/")
	fmt.Println(output)

}

func replaceWith(rawValue, target, expect string) string {
	return strings.Replace(rawValue, target, expect, 1)
}
