package main

import (
	"updev/gobasic/struct_seperate_package_pointer/employee"
)

func main() {

	// Create object from Constructor of encapsulate Customer Struct
	c := employee.InitCustomer("Karoon", "Sillapapan", 30, 20)
	c.LeavesRemaining()

	c = employee.InitCustomer("up", "Sillapapan", 30, 20) // it's cannot create new Instanse because we are implement singleton
	c.LeavesRemaining()

	/*
			Karoon Sillapapan has 10 leaves remaining
		    Karoon Sillapapan has 10 leaves remaining
	*/

}
