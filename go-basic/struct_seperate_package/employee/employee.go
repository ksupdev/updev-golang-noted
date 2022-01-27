package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

type customer struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func NewCustomer(firstName string, lastName string, totalLeaves int, leavesTaken int) customer {
	return customer{FirstName: firstName, LastName: lastName, TotalLeaves: totalLeaves, LeavesTaken: leavesTaken}
}

func (e customer) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
