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
			fmt.Printf("%.0f\t", m.M[i][x])
		}
		fmt.Println("")
	}
}

// reference: http://rosettacode.org/wiki/Reduced_row_echelon_form#2D_representation
func (m *Matrix) RowReduction() {
	lead := 0
	rowCount := m.Rows
	columnCount := m.Cols
	for r := 0; r < rowCount; r++ {
		if lead >= columnCount {
			return
		}
		i := r
		for m.M[i][lead] == 0 {
			i++
			if rowCount == i {
				i = r
				lead++
				if columnCount == lead {
					return
				}
			}
		}
		m.M[i], m.M[r] = m.M[r], m.M[i]
		f := 1 / m.M[r][lead]
		for j, _ := range m.M[r] {
			m.M[r][j] *= f
		}
		for i = 0; i < rowCount; i++ {
			if i != r {
				f = m.M[i][lead]
				for j, e := range m.M[r] {
					m.M[i][j] -= e * f
				}
			}
		}
		lead++
	}
}
