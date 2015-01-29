package search

import (
	"testing"
)

func TestBinarySearchRecursive(t *testing.T) {
	a := []float64{0, 1, 4, 5, 6, 7, 8, 9}
	returnValue := BinarySearchRecursive(a, 4, 0, len(a))
	if a[returnValue] != 4 {
		t.Error("Cannot find!")
	}
}

func TestBinarySearchIterative(t *testing.T) {
	a := []float64{0, 1, 4, 5, 6, 7, 8, 9}
	returnValue := BinarySearchIterative(a, 1)
	if a[returnValue] != 1 {
		t.Error("Cannot find!")
	}
}
