package foreach_range

import "fmt"

func DemoForeachRange() {
	fmt.Println("Hi this is demo Foreach Range")

	// Decare fix arrays
	courses := []string{"Android", "IOS", "React"}

	for index, item := range courses {
		fmt.Printf("%d Print course item %v \n", index, item)
	}

	/*
		// You can config "_" to ignore return "index" value
		for _, item := range courses {
			fmt.Printf("%d Print course item %v \n", index, item)
		}
	*/
}
