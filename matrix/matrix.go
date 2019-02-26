package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is 2D array of ints
type Matrix [][]int

// New turns a strings representation of a matrix into a 2D
// representation of the matrix
func New(matrixInput string) (Matrix, error) {
	// Turn the string into a splice of rows in string format
	split := strings.Split(matrixInput, "\n")

	matrix := make(Matrix, 0)

	// Remove odd spacing and get the size of each row in the matrix
	trimmedRow := strings.Trim(split[0], " ")
	rowAsStr := strings.Split(trimmedRow, " ")
	rowSize := len(rowAsStr)

	for _, rowStr := range split {
		if rowStr == "" {
			return [][]int{}, errors.New("Empty row")
		}

		// Remove odd spacing and turn each row string into a
		// splice of row items as strings
		trimmedRow := strings.Trim(rowStr, " ")
		rowAsStr := strings.Split(trimmedRow, " ")

		if len(rowAsStr) != rowSize {
			return [][]int{}, errors.New("Uneven row")
		}

		rowAsInt := make([]int, 0)
		// Convert each number from string into int and catch errors
		for _, numAsStr := range rowAsStr {
			numAsInt, err := strconv.Atoi(numAsStr)

			if err != nil {
				return [][]int{}, errors.New(err.Error())
			}

			rowAsInt = append(rowAsInt, numAsInt)
		}

		matrix = append(matrix, rowAsInt)
	}

	return matrix, nil
}

// Rows creates a duplicate of the matrix and returns it
func (matrix Matrix) Rows() [][]int {
	newMatrix := make(Matrix, 0)

	for _, row := range matrix {
		rowSlice := make([]int, 0)

		for _, item := range row {
			rowSlice = append(rowSlice, item)

		}

		newMatrix = append(newMatrix, rowSlice)
	}

	return newMatrix
}

// Cols turns a matrix into a slice of column slices
func (matrix Matrix) Cols() [][]int {
	colMatrix := make(Matrix, 0)

	for colInd := range matrix[0] {
		colSlice := make([]int, 0)

		for rowInd := range matrix {
			colSlice = append(colSlice, matrix[rowInd][colInd])
		}

		colMatrix = append(colMatrix, colSlice)
	}

	return colMatrix
}

// Set sets an item in the matrix, given the row, column, and new
// value, if the row and column are valid positions in the matrix
func (matrix Matrix) Set(row, col, val int) bool {
	rowLen := len(matrix[0])
	colHei := len(matrix)

	if row >= rowLen || row < 0 {
		return false
	}
	if col >= colHei || col < 0 {
		return false
	}

	matrix[row][col] = val
	return true
}
