package main

import (
	"fmt"
)

func quickSort(sizes []int, left, right int) {
	// Implement the QuickSort algorithm here
	if left < right {
		pivotIndex := partition(sizes, left, right)
		quickSort(sizes, left, pivotIndex-1)
		quickSort(sizes, pivotIndex+1, right)
	}
}

func partition(sizes []int, left, right int) int {
	pivot := sizes[right]
	i := left - 1
	for j := left; j < right; j++ {
		// > if descending order
		if sizes[j] < pivot {
			i++
			sizes[i], sizes[j] = sizes[j], sizes[i]
		}
	}
	sizes[i+1], sizes[right] = sizes[right], sizes[i+1]
	return i + 1
}

func main() {
	celestialSizes := []int{3, 5, 2, 1, 9, 5, 7, 8} // Unsorted space rock sizes
	quickSort(celestialSizes, 0, len(celestialSizes)-1)
	for _, size := range celestialSizes {
		fmt.Print(size, " ") // Output will show the sorted sizes
	}
}
