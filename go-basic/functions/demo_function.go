package functions

import (
	"fmt"
	"strconv"
)

func RunDemoFunction() {

	// test call private func
	privateFunc()

	printValueWithArgument("Hi myname is pu")
	printValueWithMultiArgumentSameType("value1", "value2")

	returnMessage := sumFunc(1, 2, "test sum value ")
	fmt.Println(returnMessage)

	firstValue, secondValue := testFuncReturnMultiValue(1, 2)
	fmt.Printf("Test  return mulitple value first value : %d, second value : %d \n", firstValue, secondValue)

	firstValueV2, secsecondValueV2 := testSwapValue(1, 2)
	fmt.Printf("Test return multiple value first value : %d, second value : %d \n", firstValueV2, secsecondValueV2)

	firstValueV3, secsecondValueV3, message := testSwapValueV2(1, 2)
	fmt.Printf("Test return multiple value %v first value : %d, second value : %d \n", message, firstValueV3, secsecondValueV3)

}

// privateFunc is Private function, This package only can use it
// the Private function, the first character of function name is lowwer
func privateFunc() {
	fmt.Println("You are running privateFunc()")

}

//PublicFunc is Public function, Every package can use it
// the Public function, the first character of function name is upper
func PublicFunc() {
	fmt.Println("You are running publicFunc() in package functions")

}

func printValueWithArgument(msg string) {
	fmt.Println("You are running printValueWithArgument = " + msg)
}

func printValueWithMultiArgumentSameType(msg1, msg2 string) {
	fmt.Println("You are running printValueWithMultiArgumentSameType msg1 = " + msg1 + ", mag2 = " + msg2)
}

func sumFunc(input1, input2 int, message string) string {
	return message + strconv.Itoa(input1+input2)
}

func testFuncReturnMultiValue(firstVaule, secondValue int) (int, int) {
	return firstVaule, secondValue
}

func testSwapValue(val1, val2 int) (reVal1, reVal2 int) {
	reVal1 = val1
	reVal2 = val2
	return reVal1, reVal2
}

func testSwapValueV2(val1, val2 int) (reVal1, reVal2 int, message string) {
	reVal1 = val1
	reVal2 = val2
	message = "process on testSwapValue"

	// If you decare the name for return value ,you will can return the empty like this code below
	return
}

//
