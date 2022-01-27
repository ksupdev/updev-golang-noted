package variable

import (
	"fmt"
	"strconv"
)

// Global variable
var count = 0

func Demo() {

	fmt.Println("---- Explicit Declaration ----")
	// Explicit Declaration (this variable has specific type)
	var tmp1 int = 0
	var strinv_tmp2 string = "hello"
	var bool_tmp3 bool = true
	fmt.Println("tmp1 = " + strconv.Itoa(tmp1))
	fmt.Println("tmp2 = " + strconv.FormatBool(bool_tmp3))
	fmt.Println("tmp3 = " + strinv_tmp2)

	// Constance value
	const const_value string = "this value cannot change"
	/*
		const_value = "try to change value"
		//show error => cannot assign to const_value (declared const)
	*/
	fmt.Println("const_value = " + const_value)

	fmt.Println("---- Implicit Declration ----")
	//Implicit Declration
	implicit_declaration_value := "defind value not specify data type"
	fmt.Println("implicit_declaration_value = " + implicit_declaration_value)

	//Implicit Declration
	fmt.Println("---- Global variable ----")
	count++
	fmt.Println("pritn global variable = " + strconv.Itoa(count))
	testCalculateGlobalVariable()

}

func testCalculateGlobalVariable() {
	// count is global variable
	count++
	fmt.Println("pritn global variable in testCalculateGlobalVariable() = " + strconv.Itoa(count))
}
