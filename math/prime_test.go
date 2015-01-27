package math

import "testing"

func TestIsPrime(t *testing.T) {
	returnValue := IsPrime(23)
	if returnValue != true {
		t.Error("Not a prime!")
	}
}
