package logic

import "testing"

func TestFunc01(t *testing.T) {
	if 1 == 1 {
		t.Log("OK")
	} else {
		t.Error("Not ok")
	}
}

func TestFunc02(t *testing.T) {
	if 1 == 1 {
		t.Log("OK")
	} else {
		t.Error("Not ok")
	}
}

func TestFunc03(t *testing.T) {
	if 1 == 1 {
		t.Log("OK")
	} else {
		t.Error("Not ok")
	}
}
