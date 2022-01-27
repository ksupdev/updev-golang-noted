package if_else_switch_case

import "fmt"

func Demo() {
	fmt.Println("Hi, This is Demo If Else and Switch case")

	fmt.Println("---- Deme if else----")
	someValue := 10
	if someValue == 10 {
		fmt.Printf("someValue == %d \n", someValue)
	} else {
		fmt.Printf("someValue == %d \n", someValue)
	}

	// Set value from other method and check condition
	if result := doCompleted(); result == "ok" {
		fmt.Println("Job has done")
	} else {
		fmt.Println("Job is doing")
	}

	fmt.Println("---- Deme swith case ----")
	fnSwithcCase()
}

func doCompleted() string {
	return "ok"
}

func fnSwithcCase() {
	index := 10

	switch index {
	case 0:
		fmt.Println("Case = 0")
		break
	case 1:
		fmt.Println("Case = 1")
		break
	case 2:
		fmt.Println("Case = 2")
		break
	default:
		fmt.Println(" something else")
		break

	}
}
