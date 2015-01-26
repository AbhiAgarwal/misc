package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	randomArray := make([]int, 20)
	for i := range randomArray {
		randomArray[i] = rand.Intn(10)
	}
	QuickSort(randomArray)
}
