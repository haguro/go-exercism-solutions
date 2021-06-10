package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a two dimensional matrix.
type Matrix [][]int

// New parses a string representation of matrix into a Matrix.
func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))
	numCells := 0
	for i, r := range rows {
		cells := strings.Fields(r)
		if i > 0 && len(cells) != numCells {
			return nil, errors.New("row size mismatch")
		}
		numCells = len(cells)
		m[i] = make([]int, len(cells))
		for j, c := range cells {
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, fmt.Errorf("failed to convert string to integer: %s", err)
			}
			m[i][j] = n
		}
	}
	return m, nil
}

// Rows returns a copy of the matrix
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(*m))
	for i, r := range *m {
		rows[i] = make([]int, len(r))
		copy(rows[i], r)
	}
	return rows
}

// Cols returns an array of the column of the matrix (i.e. the matrix transpose).
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len((*m)[0]))
	for i, r := range *m {
		for j, c := range r {
			if cols[j] == nil {
				cols[j] = make([]int, len(*m))
			}
			cols[j][i] = c
		}
	}
	return cols
}

// Set sets the value of a certain cell within the matrix.
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}
