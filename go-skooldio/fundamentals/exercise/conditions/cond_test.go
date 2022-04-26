package conditions

import (
	"testing"
)

//isOdd(n int)

func TestIsOddCaseOddValue(t *testing.T) {
	input := 1
	want := true
	op := IsOdd(input)

	if want != op {
		t.Errorf(" %v is expected but got %v", want, op)
	}
}

func TestIsOddCaseEvenValue(t *testing.T) {
	input := 2
	want := false
	op := IsOdd(input)

	if want != op {
		t.Errorf(" %v is expected but got %v", want, op)
	}
}
