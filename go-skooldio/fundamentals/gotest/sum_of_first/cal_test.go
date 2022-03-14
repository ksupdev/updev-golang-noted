package cal

import "testing"

func TestSumOfFirstThree_v1(t *testing.T) {
	given := 3
	want := 6

	get := sunOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but %d", given, want, get)
	}
}

func TestSumOfFirstThree_v2(t *testing.T) {
	given := 3
	want := 6

	get := sunOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but %d", given, want, get)
	}
}

/*
	// test all find in directory
	go test . -v

	// test specify func name
	// run both func
	go test -run TestSumOfFirstThree -v

	// run TestSumOfFirstThree_v1 func
	go test -run TestSumOfFirstThree_v1 -v

	// run TestSumOfFirstThree_v2 func
	go test -run TestSumOfFirstThree_v2 -v


*/
