package variable

import (
	"fmt"
	"strconv"
)

func Demo() {
	fmt.Println("Demo variable")
	
	// Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true

	fmt.Println("tmp1 = " + tmp1)
	fmt.Println("tmp2 = " + tmp2)
	fmt.Println("tmp3 = " + tmp3)
}
