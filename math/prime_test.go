package math

import "testing"

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestIsPrime(t *testing.T) {
	returnValue := IsPrime(5)
	if returnValue != true {
		t.Error("Not a prime!")
	}
}

func TestGetPrimeNumbers(t *testing.T) {
	returnValue := GetPrimeNumbers(5)
	currentAnswer := []int{2, 3, 5}
	if !IntArrayEquals(returnValue, currentAnswer) {
		t.Error("Generate primes are incorrect!")
	}
}
