package math

import "testing"

func TestLCM(t *testing.T) {
	returnValue := LCM(5, 3)
	if returnValue != 15 {
		t.Error("Wrong value")
	}
}
