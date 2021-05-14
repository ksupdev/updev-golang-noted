package employee

import "fmt"

type customer struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var customerInstance *customer

func InitCustomer(firstName string, lastName string, totalLeaves int, leavesTaken int) *customer {
	if customerInstance == nil {
		customerInstance = &customer{FirstName: firstName, LastName: lastName, TotalLeaves: totalLeaves, LeavesTaken: leavesTaken}
	}
	return customerInstance
}

func (e customer) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
