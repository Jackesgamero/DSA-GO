package main

import "fmt"

func findIndices(arrA, arrB []int) []int {
	n := len(arrA)
	diffMap := make(map[int]int)

	var bestI, bestJ int
	found := false

	for j := 0; j < n; j++ {
		currentDiff := arrA[j] - arrB[j]
		target := -currentDiff

		if i, exists := diffMap[target]; exists {
			if !found {
				bestI, bestJ = i, j
				found = true
			} else {
				if i < bestI || (i == bestI && j < bestJ) {
					bestI, bestJ = i, j
				}
			}
		}

		if _, exists := diffMap[currentDiff]; !exists {
			diffMap[currentDiff] = j
		}
	}
	return []int{bestI, bestJ}
}

func main() {
	arrA := []int{2, 5, 1, 4}
	arrB := []int{3, 6, 3, 2}
	indices := findIndices(arrA, arrB)
	fmt.Printf("[%d, %d]\n", indices[0], indices[1]) // Output: [2, 3]
}
