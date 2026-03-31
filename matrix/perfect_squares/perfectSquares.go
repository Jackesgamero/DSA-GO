package main

import (
	"fmt"
	"math"
)

func DiagonalTraverseAndSquares(matrix [][]int) []int {
	// TODO: Implement the function
	rows := len(matrix)
	cols := len(matrix[0])
	row, col, dir := 0, 0, 1

	path := []int{}
	for i := 0; i < rows*cols; i++ {
		path = append(path, matrix[row][col])

		if dir == 1 {
			if row == rows-1 {
				col++
				dir = -1
			} else if col == 0 {
				row++
				dir = -1
			} else {
				row++
				col--
			}
		} else {
			if col == cols-1 {
				row++
				dir = 1
			} else if row == 0 {
				col++
				dir = 1
			} else {
				row--
				col++
			}
		}
	}

	result := []int{}
	for i, num := range path {
		root := math.Sqrt(float64(num))
		if root == float64(int(root)) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	result := DiagonalTraverseAndSquares(matrix)
	fmt.Println(result) // Expected Output: [0, 8, 9]
}
