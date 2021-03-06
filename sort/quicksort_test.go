package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	randomArray := make([]int, 20)
	for i := range randomArray {
		randomArray[i] = rand.Intn(10)
	}
	randomArray = QuickSort(randomArray)
	correctArray := sort.IntSlice(randomArray)
	if !IntArrayEquals(randomArray, correctArray) {
		t.Error("Array is not correct")
	}
}
