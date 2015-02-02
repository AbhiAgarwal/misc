package math

import "testing"

func TestMatrix(t *testing.T) {
	m := Matrix{}
	err := m.Matrix(1, 1)
	if err == nil {
		m.M[0][0] = 1
	} else {
		t.Error("Cols or rows cannot be zero")
	}
}

func TestNil(t *testing.T) {
	m := Matrix{}
	if m.Nil() != true {
		t.Error("Matrix Nil function does not work")
	}
}

func TestNumElements(t *testing.T) {
	m := Matrix{}
	m.Matrix(2, 3)
	if m.NumElements() != 6 {
		t.Error("Matrix number of elements are incorrect")
	}
}

func TestGetSize(t *testing.T) {
	m := Matrix{}
	m.Matrix(2, 3)
	x, y := m.GetSize()
	if !(x == 2 && y == 3) {
		t.Error("Matrix sizes are incorrect")
	}
}

func compare(X, Y [][]float64) bool {
	for i := 0; i < len(X); i++ {
		for x := 0; x < len(X[0]); x++ {
			if X[i][x] != Y[i][x] {
				return false
			}
		}
	}
	return true
}

func TestRowReduction(t *testing.T) {
	m := Matrix{}
	m.Matrix(3, 4)
	m.M = [][]float64{
		{1, 2, -1, -4},
		{2, 3, -1, -11},
		{-2, 0, -3, 22},
	}
	m.RowReduction()

	correctMatrix := Matrix{}
	correctMatrix.Matrix(3, 4)
	correctMatrix.M = [][]float64{
		{1, 0, 0, -8},
		{-0, 1, 0, 1},
		{-0, -0, 1, -2},
	}

	if !compare(correctMatrix.M, m.M) {
		t.Error("Row reduction not working!")
	}
}
