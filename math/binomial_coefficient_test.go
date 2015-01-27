package math

import "testing"

func TestBinomialCoefficient(t *testing.T) {
	returnValue := BinomialCoefficient(5, 3)
	if returnValue != 10 {
		t.Error("Wrong value")
	}
}
