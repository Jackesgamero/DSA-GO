package main

import (
	"fmt"
)

func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	left, right := append([]int{}, arr[:mid]...), append([]int{}, arr[mid:]...)

	mergeSort(left)
	mergeSort(right)
	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	i, j := 0, 0
	for k := 0; k < len(arr); k++ {
		if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}

func main() {
	musicLibrary := []int{3, 5, 1, 2, 6, 9, 8, 8, 3}
	mergeSort(musicLibrary)

	for _, songID := range musicLibrary {
		fmt.Print(songID, " ")
	}
}
