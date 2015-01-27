package main

import "fmt"

type FenwickTree struct {
	Ft []int
}

func NewFenwickTree() *FenwickTree {
	return &FenwickTree{}
}

func (f *FenwickTree) FenwickTree(n int) {
	f.Ft = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		f.Ft[i] = 0
	}
}

func (f *FenwickTree) rsqI(b int) int {
	var sum int = 0
	for ; b > 0; b -= (b & (-b)) {
		sum += f.Ft[b]
	}
	return sum
}

func (f *FenwickTree) Rsq(a, b int) int {
	returnValue := f.rsqI(b)
	if a == 1 {
		return returnValue
	} else {
		return returnValue - f.rsqI(a-1)
	}
}

func (f *FenwickTree) Adjust(k, v int) {
	for ; k < len(f.Ft); k += (k & (-k)) {
		f.Ft[k] += v
	}
}

func main() {
	fenwichTree := NewFenwickTree()
	fenwichTree.FenwickTree(5)
	fenwichTree.Adjust(5, 2)
	fmt.Println(fenwichTree.Rsq(5, 2))
}
