package math

import "testing"

func TestGCD(t *testing.T) {
	returnValue := GCD(50, 40)
	if returnValue != 10 {
		t.Error("Wrong value")
	}
}
