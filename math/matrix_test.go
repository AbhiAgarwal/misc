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
