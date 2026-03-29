package main

import "fmt"

func solution(matrix [][]int) [][]int {
	// TODO: Your solution here
	rows := len(matrix)
	cols := len(matrix[0])
	result := [][]int{}

	row, col, dir := 0, 0, -1
	for i := 0; i < rows*cols; i++ {
		if matrix[row][col] < 0 {
			result = append(result, []int{row, col})
		}

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

	return result
}

func main() {
	exampleMatrix := [][]int{
		{1, -2, 3, -4},
		{5, -6, 7, 8},
		{-9, 10, -11, 12},
	}
	result := solution(exampleMatrix)
	for _, indexPair := range result {
		fmt.Printf("[%d, %d]\n", indexPair[0], indexPair[1])
	}
}
