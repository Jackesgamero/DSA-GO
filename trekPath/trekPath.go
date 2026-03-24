package main

import (
	"fmt"
)

func TrekPath(elevationMap [][]int, startRow int, startCol int) []int {
	// Initialize the path with the starting position's elevation.
	path := []int{elevationMap[startRow][startCol]}
	// Implement the loop to find the path through higher elevations.
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	bestHeight := elevationMap[startRow][startCol]
	bestRow, bestCol := -1, -1

	for _, dir := range directions {
		r := startRow + dir[0]
		c := startCol + dir[1]

		if (r >= 0 && r < len(elevationMap)) &&
			(c >= 0 && c < len(elevationMap[0])) &&
			elevationMap[r][c] > bestHeight {
			bestHeight = elevationMap[r][c]
			bestRow, bestCol = r, c
		}
	}

	if bestRow == -1 {
		return path
	}

	nextPath := TrekPath(elevationMap, bestRow, bestCol)

	return append(path, nextPath...)
}

func main() {
	mountain := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 5, 6},
	}
	result := TrekPath(mountain, 1, 1)
	for _, height := range result {
		fmt.Printf("%d ", height)
	}
}
