package math

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Rows int
	Cols int
	M    [][]float64
}

type error interface {
	Error() string
}

func (m *Matrix) Matrix(rows, cols int) error {
	if rows == 0 || cols == 0 {
		return errors.New("Values can not be negative")
	}
	m.Rows = rows
	m.Cols = cols
	m.setMatrix()
	return nil
}

func (m *Matrix) setMatrix() {
	m.M = make([][]float64, m.Rows)
	for i := range m.M {
		m.M[i] = make([]float64, m.Cols)
	}
}

func (m *Matrix) Nil() bool {
	return m.Rows == 0 || m.Cols == 0
}

func (m *Matrix) NumElements() int {
	return m.Rows * m.Cols
}

func (m *Matrix) GetSize() (rows, cols int) {
	rows = m.Rows
	cols = m.Cols
	return
}

func (m *Matrix) PrintMatrix() {
	for i := 0; i < m.Rows; i++ {
		for x := 0; x < m.Cols; x++ {
			fmt.Printf("%.0f ", m.M[i][x])
		}
		fmt.Println("")
	}
}
