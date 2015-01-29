package datastructure

import (
	"testing"
)

func TestFenwickTree(t *testing.T) {
	n := 3
	fenwichtree := NewFenwickTree()
	fenwichtree.FenwickTree(n)
	if len(fenwichtree.Ft) != n+1 {
		t.Error("Wrong size!")
	}
}
