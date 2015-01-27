package sort

import (
	"math/rand"
	"sort"
	"testing"
)

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

func TestMergeSort(t *testing.T) {
	randomArray := make([]int, 20)
	for i := range randomArray {
		randomArray[i] = rand.Intn(10)
	}
	randomArray = MergeSort(randomArray)
	correctArray := sort.IntSlice(randomArray)
	if !IntArrayEquals(randomArray, correctArray) {
		t.Error("Array is not correct")
	}
}
