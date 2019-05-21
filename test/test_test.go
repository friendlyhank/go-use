package test

import "testing"

var b = 0

func TestDivision(t *testing.T) {
	if b == 0 {
		t.Error("B不能为零")
	} else {
		t.Log("B不为零")
	}
}
