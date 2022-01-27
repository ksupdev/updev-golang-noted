package main

import (
	"updev/gobasic/struct_seperate_package/employee"
)

func main() {
	// Create object from Employee Struct
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	// Create object from Constructor of encapsulate Customer Struct
	c := employee.NewCustomer("Karoon", "Sillapapan", 30, 20)
	c.LeavesRemaining()

}
